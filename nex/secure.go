package nex

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	nex "github.com/PretendoNetwork/nex-go/v2"
	common_ranking "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	common_utility "github.com/PretendoNetwork/nex-protocols-common-go/v2/utility"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
	"github.com/Protarium-Network/zen-pinball-2-nex/database"
	"github.com/Protarium-Network/zen-pinball-2-nex/globals"
)

var nexUniqueIDCounter atomic.Uint64

func init() { nexUniqueIDCounter.Store(uint64(time.Now().Unix()) << 16) }

var SecureServer *nex.PRUDPServer
var SecureEndpoint *nex.PRUDPEndPoint

func StartSecureServer() {
	SecureServer = nex.NewPRUDPServer()

	// See authentication.go.
	SecureServer.PRUDPV1Settings.LegacyConnectionSignature = globals.LegacyConnectionSignature

	SecureEndpoint = nex.NewPRUDPEndPoint(1)
	SecureEndpoint.IsSecureEndPoint = true
	SecureEndpoint.ServerAccount = globals.SecureServerAccount
	SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	SecureServer.BindPRUDPEndPoint(SecureEndpoint)

	SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(int(globals.NEXMajor), int(globals.NEXMinor), int(globals.NEXPatch)))
	SecureServer.AccessKey = globals.AccessKey
	// See authentication.go.
	SecureServer.ByteStreamSettings.UseStructureHeader = globals.UseStructureHeader

	SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()
		if request == nil {
			return
		}
		pid := uint64(packet.Sender().PID())
		fmt.Printf("[ZP2 Secure] PID=%d protocol=0x%02X method=0x%02X\n", pid, request.ProtocolID, request.MethodID)
	})

	SecureEndpoint.OnConnectionEnded(func(connection *nex.PRUDPConnection) {
		fmt.Printf("[ZP2 Secure] PID=%d disconnected\n", uint64(connection.PID()))
	})

	registerSecureServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_ZP2_SECURE_PORT"))
	globals.Logger.Successf("[ZP2] Secure server listening on UDP %d", port)
	SecureServer.Listen(port)
}

// registerSecureServerProtocols: the protocols the game links (recovered from
// jsextension_ext-nex-gameserver.rpl): SecureConnection, Utility, Ranking.
// No DataStore, no MatchMaking in the client.
func registerSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	secureCommon := common_secure.NewCommonProtocol(secureProtocol)
	secureCommon.EnableInsecureRegister()
	secureCommon.CreateReportDBRecord = database.CreateReportDBRecord

	utilityProtocol := utility.NewProtocol()
	SecureEndpoint.RegisterServiceProtocol(utilityProtocol)
	utilityCommon := common_utility.NewCommonProtocol(utilityProtocol)
	// The game calls AcquireNexUniqueID before uploading scores; without a
	// generator the call fails and UploadScore is never sent. Seeded from the
	// clock so IDs stay unique across restarts.
	utilityCommon.GenerateNEXUniqueID = func() uint64 { return nexUniqueIDCounter.Add(1) }

	rankingProtocol := ranking.NewProtocol()
	SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	rankingCommon := common_ranking.NewCommonProtocol(rankingProtocol)
	rankingCommon.GetRankingsAndCountByCategoryAndRankingOrderParam = database.ZP2GetRankingsAndCountByCategoryAndRankingOrderParam
	rankingCommon.GetRankingsByMode = database.ZP2GetRankings
	rankingCommon.GetCommonData = database.ZP2GetCommonData
	rankingCommon.UploadCommonData = database.ZP2UploadCommonData
	rankingCommon.InsertRankingByPIDAndRankingScoreData = database.ZP2InsertRankingByPIDAndRankingScoreData
}

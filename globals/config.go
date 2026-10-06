package globals

import (
	"fmt"
	"os"
	"strings"
)

// NEX configuration for "Zen Pinball 2" (Wii U, Zen Studios, 2013).
//
// Game server ID 0x10113800 / access key come from kinnay's nexwiiu.json
// (captured from the retail game, build 3_4_13_3_0, branch release/ngs/3.4.x.3).
// The retail binary links only Ranking, Utility, SecureConnection and
// AccountManagement (see RECON.md).
const (
	GameServerID = "10113800" // 0x10113800
	AccessKey    = "2ff15f7e"
)

// NEX 3.4.13 is a "3.4.x" build: legacy connection signature on, no
// structure header (same as Mario & Sonic Sochi / Trine 2). Env overrides:
//
//	PN_ZP2_NEX_VERSION=3.4.13
//	PN_ZP2_LEGACY_SIGNATURE=0|1
//	PN_ZP2_STRUCTURE_HEADER=0|1
var (
	NEXMajor, NEXMinor, NEXPatch uint32 = 3, 4, 13

	LegacyConnectionSignature = true
	UseStructureHeader        = false
)

func init() {
	if v := strings.TrimSpace(os.Getenv("PN_ZP2_NEX_VERSION")); v != "" {
		var a, b, c uint32
		if n, _ := fmt.Sscanf(v, "%d.%d.%d", &a, &b, &c); n == 3 {
			NEXMajor, NEXMinor, NEXPatch = a, b, c
		}
	}
	if v := os.Getenv("PN_ZP2_LEGACY_SIGNATURE"); v != "" {
		LegacyConnectionSignature = v == "1"
	}
	if v := os.Getenv("PN_ZP2_STRUCTURE_HEADER"); v != "" {
		UseStructureHeader = v == "1"
	}
}

# Recon - Zen Pinball 2 (Wii U)

Source: `t.rpx` (build path `Y:\pinball\YAP_wiiu	.rpx`), zlib sections
decompressed with `recon/unrpx.py`, plus kinnay's `nexwiiu.json`.

- Kinnay: game server ID `0x10113800` (title `0005000010113800`), key
  `2ff15f7e`, `branch:origin/release/ngs/3.4.x.3`, build `3_4_13_3_0`.
  The sibling entry Star Wars Pinball (`0x10132a00`, key `ecd0e530`) shares the
  engine and could be added the same way.
- NEX is statically linked (`nn::nex::*`), SDK tags `NEX_3_0_1_4`,
  `NEX_UT_3_0_1`, `NEX_RK_3_0_1`. Linked protocols: Ranking, Utility,
  SecureConnection, AccountManagement, TicketGranting. No DataStore, no
  MatchMaking.
- Game side: `PlatformLib::WIIUNEXSystem` (`NEX_SetScore`, weekly scores are
  ignored client-side), `cPinballSocketServer`, `RankingScoreData`,
  `RankingOrderParam`, `RankingChangeAttributesParam`, `RankingStats`.
- `AcquireNexServiceToken` log: `GameID:%08x Success:%d ...`.
- Access key is not a literal in the binary (loaded from game config);
  taken from kinnay.
- Region IDs: only the 0x10113800 ID is documented; other regions are
  unverified (see README).

`recon/*.bin` are derived from the proprietary executable and git-ignored.

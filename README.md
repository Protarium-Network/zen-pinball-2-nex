# Zen Pinball 2 (Wii U) - NEX server

Preservation NEX server for **Zen Pinball 2** (Zen Studios, 2013),
`game_server_id` `10113800`. Built on the
[Pretendo Network](https://github.com/PretendoNetwork) NEX libraries, same
layout as the other Protarium Wii U servers.

## Status

Builds and unit-tests clean; **not yet console-tested**.

Scope: Ticket Granting, Secure Connection, Utility and **Ranking** (table
leaderboards) - exactly what the game's statically linked NEX client uses. No
DataStore, no matchmaking. See [PROTOCOL_COVERAGE.md](PROTOCOL_COVERAGE.md).

## Recovered configuration

| Field | Value | Source |
|---|---|---|
| Game server ID | `10113800` | kinnay's `nexwiiu.json` |
| Access key | `2ff15f7e` | kinnay's `nexwiiu.json` |
| NEX version | `3.4.13` (`3_4_13_3_0`) | kinnay's `nexwiiu.json` |
| `LegacyConnectionSignature` | `true` (3.4.x) | `PN_ZP2_LEGACY_SIGNATURE=0` to flip |
| `UseStructureHeader` | `false` (3.4.x) | `PN_ZP2_STRUCTURE_HEADER=1` to flip |

If a console loops on CONNECT (`106-0502`) flip the signature switch. See
[RECON.md](RECON.md).

## Running

```bash
cp .env.example .env
cp settings.example.json settings.json
docker compose up --build        # UDP 28500 (auth) / 28501 (secure)
```

Local mode (`PN_ZP2_LOCAL_MODE=1`, default) takes accounts from
`settings.json` and does not validate tokens — isolated networks only.
**Shared mode:** `PN_ZP2_LOCAL_MODE=0` plus `PN_ZP2_NEX_TOKEN_AES_KEY` (64 hex)
and `PN_ZP2_NEX_PASSWORD_SECRET` (≥32 bytes hex) matching your account server.
`PN_ZP2_SECURE_HOST` must be short (~15 chars).

## License

AGPL-3.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).

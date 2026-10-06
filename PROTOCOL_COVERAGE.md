# Protocol coverage

| Endpoint | Protocol | Coverage |
|---|---|---|
| Auth | Ticket Granting | Login / LoginEx / RequestTicket (stock) |
| Secure | Secure Connection, Utility | stock |
| Secure | Ranking | GetRanking, GetCommonData, UploadScore, Upload/DeleteCommonData; Postgres `zp2_*` tables |

Not registered: DataStore, NAT Traversal, MatchMaking*. The game's NEX
code links Account Management, Authentication, Ranking, Secure Connection and
Utility only (RECON.md).

Unverified on a console.

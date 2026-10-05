# Forum: Full Stack Community

Mandatory · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 4 weeks after preparation

Integrate the prepared HTTP, HTML, SQLite, credential, session, reaction, filter and Docker components into one usable forum for public readers and registered contributors.

## Preparation gate

- P01: Node Runtime and Package Lab
- P02: HTTP Message and Bounded Body Lab
- P03: HTML Forms and Safe Rendering Board
- P04: Community ERD and Relational Schema Design
- P05: SQLite Schema and Migration Workbench
- P06: Topic and Reply Repository
- P07: Reaction State Machine and Transactions
- P08: Feed Filters, Aggregates and Pagination
- P09: Registration and Password Hashing Vault
- P10: Session Cookies, Route Guards and CSRF
- P11: Failure Handling, Tests and Team Integration
- P12: Docker Images, Volumes and Restore Drill

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **B01** Use SQLite persistence with directly written CREATE, INSERT and SELECT queries; data survives application and container replacement. Prep: P05, P06, P12. Evidence: Inspect SQL and reproduce restart persistence. (Source requirement.)
- **B02** Submit an ERD, relational schema, data dictionary, key constraints and deletion policy. Prep: P04, P05. Evidence: Peer maps every stored feature to a reviewed relation. (Academy: promoted from source recommendation.)
- **B03** Registration requires email, username and password; normalize under P09 policy; reject duplicate normalized email and username with 409 and invalid fields with 422. Prep: P09. Evidence: Race duplicate registration: one member and one conflict. (Source requirement.)
- **B04** Hash every password with the prepared versioned scrypt record or a preapproved equivalent; never store or log plaintext credentials. Prep: P09. Evidence: Inspect DB/logs and verify correct and incorrect credentials. (Academy: source bonus promoted to baseline.)
- **B05** Login verifies credentials, issues an opaque cookie session, permits only one active session per member, enforces an 8-hour server expiry and supports POST logout. Prep: P10. Evidence: Second browser login revokes first; copied expired/revoked tokens fail. (Source requirement.)
- **B06** Only authenticated members may create posts and comments; guests may read public posts and comments. Author identity comes from the session. Prep: P06, P10. Evidence: Direct guest writes fail without persisted changes. (Source requirement.)
- **B07** A post has a nonblank 1–120-code-point title, 1–5000-code-point body and 1–5 distinct existing categories; comment body is 1–5000. Creation and category linkage are atomic. Prep: P03, P06. Evidence: An invalid final category produces no partial post. (Source categories plus academy field bounds.)
- **B08** Authenticated members may like or dislike both posts and comments. One current reaction per member-target; repeat set is idempotent, opposite set switches, clear removes. Prep: P07, P10. Evidence: All nine state transitions and both target types pass. (Source requirement.)
- **B09** Every visitor sees correct like and dislike counts, including zeros. Multiple categories/comments must not multiply counts. Prep: P07, P08. Evidence: Two-category, three-comment fixture has exact hand-calculated counts. (Source requirement.)
- **B10** Filter posts by category, the logged-in member’s created posts, and the logged-in member’s currently liked posts. Combined filters use AND; private filters require login. Prep: P08, P10. Evidence: Guest mine/liked requests fail; client-supplied member ID is ignored. (Source requirement.)
- **B11** Use native HTML and server rendering without React, Angular, Vue or other frontend frameworks. Provide labelled forms, usable empty/error pages and safe escaped text. Prep: P03. Evidence: Keyboard workflow and hostile-text fixture pass. (Source requirement.)
- **B12** Handle method/path/input/domain/server/storage errors with the prepared status policy; never expose SQL, stack traces, passwords or private paths. Prep: P02, P11. Evidence: Malformed requests and injected faults leave the server usable. (Source requirement.)
- **B13** Use Docker with locked dependencies, supported platform documentation, a non-root application process and persistent database directory. Prep: P01, P12. Evidence: A teammate builds and replaces the container from the README. (Source requirement.)
- **B14** Protect unsafe forms with CSRF and route authorization; private responses use an explicit cache policy. Maintain isolated unit and integration checks. Prep: P10, P11. Evidence: Forged form, forged member ID and failed transaction cause no write. (Academy baseline safety and verification.)
- **B15** Document modular contracts, configuration, seed data, backup/restore and each team member’s reviewed contribution. Prep: P01, P11, P12. Evidence: Every member can change and explain a component they did not originally author. (Academy assessment policy.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- A team-owned Node.js or Go repository with a working browser forum
- ERD, schema/migrations, automated tests, Docker files, operations README and an audit demonstration

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **Registration race:** same normalized email registered concurrently → one account; other request 409.
- **Session replacement:** two browser logins for one member → only latest active token usable.
- **Public access:** guest opens post with comments and counts → 200; correct public content.
- **Write protection:** guest or forged author submits post/comment/reaction → 401/403; no write.
- **Atomic post:** categories [valid,missing] → validation/not-found response; no partial post.
- **Post/comment reactions:** repeat like, switch dislike, clear on each target → exact counts; one current reaction.
- **Filters:** category+liked; mine from another memberId param → correct AND result; member comes from session.
- **Hostile input:** HTML and SQL-like text → safe literal rendering; schema intact.
- **Fault recovery:** malformed body and injected storage fault → safe status/message; next health request succeeds.
- **Container persistence:** build, create data, replace container → data unchanged; non-root process; successful separate restore.

## Full guided chapters

[Open the HTML guide](../projects/f01-forum-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../forum.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum). Source Go restrictions are replaced explicitly by this academy language choice.

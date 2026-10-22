# Forum: Full Stack Community — Go

Mandatory · Go learning path · 4–5 students · 4 weeks after preparation

Integrate the prepared HTTP, HTML, SQLite, credential, session, reaction, filter and Docker components into one usable forum for public readers and registered contributors.

## Preparation gate

- [P01: Go Language, Runtime and Module Lab](../prep/p01-go-runtime-lab-guide/index.html)
- [P02: HTTP Message and Bounded Body Lab](../prep/p02-http-message-lab-guide/index.html)
- [P03: HTML Forms and Safe Rendering Board](../prep/p03-html-form-board-guide/index.html)
- [P04: Community ERD and Relational Schema Design](../prep/p04-community-erd-design-guide/index.html)
- [P05: SQLite Schema and Migration Workbench](../prep/p05-sqlite-migration-workbench-guide/index.html)
- [P06: Topic and Reply Repository](../prep/p06-topic-repository-guide/index.html)
- [P07: Reaction State Machine and Transactions](../prep/p07-reaction-state-machine-guide/index.html)
- [P08: Feed Filters, Aggregates and Pagination](../prep/p08-feed-filter-engine-guide/index.html)
- [P09: Registration and Password Hashing Vault](../prep/p09-credential-vault-guide/index.html)
- [P10: Session Cookies, Route Guards and CSRF](../prep/p10-session-cookie-gateway-guide/index.html)
- [P11: Failure Handling, Tests and Team Integration](../prep/p11-failure-and-team-lab-guide/index.html)
- [P12: Docker Images, Volumes and Restore Drill](../prep/p12-docker-persistence-lab-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **B01** Use SQLite persistence with directly written CREATE, INSERT and SELECT queries; data survives application and container replacement. Prep: P05, P06, P12. Evidence: Inspect SQL and reproduce restart persistence.
- **B02** Submit an ERD, relational schema, data dictionary, key constraints and deletion policy. Prep: P04, P05. Evidence: Peer maps every stored feature to a reviewed relation.
- **B03** Registration requires email, username and password; normalize under P09 policy; reject duplicate normalized email and username with 409 and invalid fields with 422. Prep: P09. Evidence: Race duplicate registration: one member and one conflict.
- **B04** Hash every password with the prepared versioned scrypt record or a preapproved equivalent; never store or log plaintext credentials. Prep: P09. Evidence: Inspect DB/logs and verify correct and incorrect credentials.
- **B05** Login verifies credentials, issues an opaque cookie session, permits only one active session per member, enforces an 8-hour server expiry and supports POST logout. Prep: P10. Evidence: Second browser login revokes first; copied expired/revoked tokens fail.
- **B06** Only authenticated members may create posts and comments; guests may read public posts and comments. Author identity comes from the session. Prep: P06, P10. Evidence: Direct guest writes fail without persisted changes.
- **B07** A post has a nonblank 1–120-code-point title, 1–5000-code-point body and 1–5 distinct existing categories; comment body is 1–5000. Creation and category linkage are atomic. Prep: P03, P06. Evidence: An invalid final category produces no partial post.
- **B08** Authenticated members may like or dislike both posts and comments. One current reaction per member-target; repeat set is idempotent, opposite set switches, clear removes. Prep: P07, P10. Evidence: All nine state transitions and both target types pass.
- **B09** Every visitor sees correct like and dislike counts, including zeros. Multiple categories/comments must not multiply counts. Prep: P07, P08. Evidence: Two-category, three-comment fixture has exact hand-calculated counts.
- **B10** Filter posts by category, the logged-in member’s created posts, and the logged-in member’s currently liked posts. Combined filters use AND; private filters require login. Prep: P08, P10. Evidence: Guest mine/liked requests fail; client-supplied member ID is ignored.
- **B11** Use native HTML and server rendering without React, Angular, Vue or other frontend frameworks. Provide labelled forms, usable empty/error pages and safe escaped text. Prep: P03. Evidence: Keyboard workflow and hostile-text fixture pass.
- **B12** Handle method/path/input/domain/server/storage errors with the prepared status policy; never expose SQL, stack traces, passwords or private paths. Prep: P02, P11. Evidence: Malformed requests and injected faults leave the server usable.
- **B13** Use Docker with locked dependencies, supported platform documentation, a non-root application process and persistent database directory. Prep: P01, P12. Evidence: A teammate builds and replaces the container from the README.
- **B14** Protect unsafe forms with CSRF and route authorization; private responses use an explicit cache policy. Maintain isolated unit and integration checks. Prep: P10, P11. Evidence: Forged form, forged member ID and failed transaction cause no write.
- **B15** Document modular contracts, configuration, seed data, backup/restore and each team member’s reviewed contribution. Prep: P01, P11, P12. Evidence: Every member can change and explain a component they did not originally author.

## Submission

- A team-owned Node.js or Go repository with a working browser forum
- ERD, schema/migrations, automated tests, Docker files, operations README and an audit demonstration

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f01-forum-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

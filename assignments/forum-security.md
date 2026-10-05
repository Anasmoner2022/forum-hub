# Forum Security: HTTPS, Limits and Session Defense

Optional · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 1 week after preparation

Harden the running forum using the prepared certificate, rate-limit, password and server-session defenses, then demonstrate them under hostile and concurrent requests.

## Preparation gate

- F01: Forum: Full Stack Community
- P18: HTTPS, Certificates and Container Configuration
- P19: Rate Limiting and Bounded Work
- P20: Session Attack and Browser Defense Workbench

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **S01** Serve HTTPS with a certificate and private key; development trust is explicit and deployed trust uses a valid CA-issued hostname certificate with renewal documentation. Prep: P18. Evidence: Verified HTTPS request and wrong-hostname failure. (Source requirement.)
- **S02** Document certificate authority, hostname/SAN verification, TLS protocol floor and negotiated cipher concepts; use reviewed runtime defaults and minimum TLSv1.2. Prep: P18. Evidence: Inspect certificate and protocol without globally bypassing verification. (Source cipher recommendation made an academy explanation task.)
- **S03** Enforce login identifier/IP and authenticated-write rate policies from P19; return 429 with Retry-After and bound costly concurrent work and limiter key storage. Prep: P19. Evidence: Budget, boundary, spoofed-header and work-count tests pass. (Source requirement.)
- **S04** Retain measured versioned password hashing and unique salted records; passwords are hashed, not reversibly encrypted. Prep: P09, P19, P20. Evidence: Correct/incorrect verification and plaintext/log scans pass. (Source requirement.)
- **S05** Keep session state server-side with unique random opaque cookies, digests, one active session, idle/absolute expiration, logout revocation and privilege-change rotation. Prep: P10, P20. Evidence: Fixation, forged/revoked token and simultaneous-login tests pass. (Source requirement.)
- **S06** Protect unsafe routes with session/CSRF/origin rules, escape rendered text, bind SQL values and deploy reviewed safe headers/private cache policy. Prep: P03, P06, P10, P20. Evidence: Forged form and hostile-content tests produce no unauthorized mutation. (Academy defense detail.)
- **S07** Handle security/technical failures without exposing secrets and test pressure against a disposable instance; keep keys out of source/image layers. Prep: P11, P18, P20. Evidence: Injected faults and log scans pass; image excludes private keys. (Source requirement.)
- **S08** Retain the F01 suite and, for installed optional features, re-run their callback/upload/visibility guards under HTTPS and rate limits. Prep: P11, P20. Evidence: Every installed workflow passes hardened-mode regression. (Academy integrated regression detail.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- HTTPS Docker configuration and route-limit policy
- Threat model, attack regression suite and certificate/secret operations procedure

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **HTTPS:** verified localhost request with --cacert → 200; SAN/time verified.
- **Bad configuration:** missing private key → startup refuses before listen.
- **Login limits:** sixth identifier attempt inside 60 seconds → 429 with Retry-After.
- **Work concurrency:** burst across many identifiers → configured active hash bound.
- **Session fixation:** chosen pre-login token then successful login → fresh random token; chosen token fails.
- **Revocation/expiry:** logout replay and idle/absolute clock boundaries → no authentication.
- **CSRF/injection:** hostile form/text/SQL-like payload → no unauthorized write or executable content.
- **Regression:** all installed forum features over HTTPS → same permitted behavior; no secret leakage.

## Full guided chapters

[Open the HTML guide](../projects/f04-forum-security-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../forum-security.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum/security). Source Go restrictions are replaced explicitly by this academy language choice.

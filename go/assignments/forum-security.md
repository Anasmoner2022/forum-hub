# Forum Security: HTTPS, Limits and Session Defense — Go

Optional · Go learning path · 4–5 students · 1 week after preparation

Harden the running forum using the prepared certificate, rate-limit, password and server-session defenses, then demonstrate them under hostile and concurrent requests.

## Preparation gate

- [F01: Forum: Full Stack Community](../projects/f01-forum-guide/index.html)
- [P18: HTTPS, Certificates and Container Configuration](../prep/p18-https-certificate-lab-guide/index.html)
- [P19: Rate Limiting and Bounded Work](../prep/p19-rate-limit-service-guide/index.html)
- [P20: Session Attack and Browser Defense Workbench](../prep/p20-session-attack-workbench-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **S01** Serve HTTPS with a certificate and private key; development trust is explicit and deployed trust uses a valid CA-issued hostname certificate with renewal documentation. Prep: P18. Evidence: Verified HTTPS request and wrong-hostname failure.
- **S02** Document certificate authority, hostname/SAN verification, TLS protocol floor and negotiated cipher concepts; use reviewed runtime defaults and minimum TLSv1.2. Prep: P18. Evidence: Inspect certificate and protocol without globally bypassing verification.
- **S03** Enforce login identifier/IP and authenticated-write rate policies from P19; return 429 with Retry-After and bound costly concurrent work and limiter key storage. Prep: P19. Evidence: Budget, boundary, spoofed-header and work-count tests pass.
- **S04** Retain measured versioned password hashing and unique salted records; passwords are hashed, not reversibly encrypted. Prep: P09, P19, P20. Evidence: Correct/incorrect verification and plaintext/log scans pass.
- **S05** Keep session state server-side with unique random opaque cookies, digests, one active session, idle/absolute expiration, logout revocation and privilege-change rotation. Prep: P10, P20. Evidence: Fixation, forged/revoked token and simultaneous-login tests pass.
- **S06** Protect unsafe routes with session/CSRF/origin rules, escape rendered text, bind SQL values and deploy reviewed safe headers/private cache policy. Prep: P03, P06, P10, P20. Evidence: Forged form and hostile-content tests produce no unauthorized mutation.
- **S07** Handle security/technical failures without exposing secrets and test pressure against a disposable instance; keep keys out of source/image layers. Prep: P11, P18, P20. Evidence: Injected faults and log scans pass; image excludes private keys.
- **S08** Retain the F01 suite and, for installed optional features, re-run their callback/upload/visibility guards under HTTPS and rate limits. Prep: P11, P20. Evidence: Every installed workflow passes hardened-mode regression.

## Submission

- HTTPS Docker configuration and route-limit policy
- Threat model, attack regression suite and certificate/secret operations procedure

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f04-forum-security-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

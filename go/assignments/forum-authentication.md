# Forum Authentication: Google and GitHub SSO — Go

Optional · Go learning path · 4–5 students · 1 week after preparation

Add Google and GitHub registration/login to the forum while preserving one local account, one session policy and the prepared identity trust rules.

## Preparation gate

- [F01: Forum: Full Stack Community](../projects/f01-forum-guide/index.html)
- [P13: External API Client and Secret Handling](../prep/p13-external-api-client-guide/index.html)
- [P14: OAuth Authorization Code, State and PKCE Lab](../prep/p14-oauth-callback-lab-guide/index.html)
- [P15: Google OIDC, GitHub Identity and Account Linking](../prep/p15-provider-identity-lab-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **A01** Support both Google and GitHub for first-time registration and returning-user login. Local credential login remains usable. Prep: P15. Evidence: Live demo with each provider plus local login.
- **A02** Use authorization-code flow with browser-bound, five-minute, single-use state and S256 PKCE; exact callback URIs and provider errors are handled. Prep: P14. Evidence: Replay, cross-browser state, expired attempt and wrong verifier all fail.
- **A03** Verify Google OIDC signature, issuer, audience, expiry, algorithm policy and nonce through maintained verification code and trusted discovery/JWKS. Prep: P15. Evidence: Synthetic wrong-claim tests create no account/session.
- **A04** Fetch authenticated GitHub identity server-side; key it by stable provider user ID and handle private/missing verified email through explicit onboarding. Prep: P13, P15. Evidence: Renamed login retains account; missing email does not invent identity.
- **A05** Persist unique provider+subject to member mappings; repeated/concurrent login does not duplicate accounts. Matching email alone never links an existing member. Prep: P09, P15. Evidence: Email collision cannot take over a local account.
- **A06** If offering provider linking, require current account proof, recent reauthentication and a browser-bound link intent; reject a provider identity linked elsewhere. Prep: P15. Evidence: Cross-account link conflict leaves original mapping intact.
- **A07** Keep secrets/tokens server-side, bound outbound calls and sanitize errors; do not use provider tokens as forum session cookies. Prep: P13, P15. Evidence: Timeout/oversize scenarios fail safely; secret-marker scan is clean.
- **A08** Issue the existing forum session after identity resolution, enforce old-session invalidation and preserve all F01 guards and tests. Prep: P10, P15. Evidence: SSO login in browser B revokes browser A; base regression passes.

## Submission

- Two real SSO provider flows integrated into the forum
- Provider setup documentation, identity migration and callback/security regression suite

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f02-forum-authentication-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

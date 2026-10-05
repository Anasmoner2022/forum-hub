# Forum Authentication: Google and GitHub SSO

Optional · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 1 week after preparation

Add Google and GitHub registration/login to the forum while preserving one local account, one session policy and the prepared identity trust rules.

## Preparation gate

- F01: Forum: Full Stack Community
- P13: External API Client and Secret Handling
- P14: OAuth Authorization Code, State and PKCE Lab
- P15: Google OIDC, GitHub Identity and Account Linking

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **A01** Support both Google and GitHub for first-time registration and returning-user login. Local credential login remains usable. Prep: P15. Evidence: Live demo with each provider plus local login. (Source requirement.)
- **A02** Use authorization-code flow with browser-bound, five-minute, single-use state and S256 PKCE; exact callback URIs and provider errors are handled. Prep: P14. Evidence: Replay, cross-browser state, expired attempt and wrong verifier all fail. (Academy security detail.)
- **A03** Verify Google OIDC signature, issuer, audience, expiry, algorithm policy and nonce through maintained verification code and trusted discovery/JWKS. Prep: P15. Evidence: Synthetic wrong-claim tests create no account/session. (Academy identity verification detail.)
- **A04** Fetch authenticated GitHub identity server-side; key it by stable provider user ID and handle private/missing verified email through explicit onboarding. Prep: P13, P15. Evidence: Renamed login retains account; missing email does not invent identity. (Academy identity resolution detail.)
- **A05** Persist unique provider+subject to member mappings; repeated/concurrent login does not duplicate accounts. Matching email alone never links an existing member. Prep: P09, P15. Evidence: Email collision cannot take over a local account. (Academy identity integrity detail.)
- **A06** If offering provider linking, require current account proof, recent reauthentication and a browser-bound link intent; reject a provider identity linked elsewhere. Prep: P15. Evidence: Cross-account link conflict leaves original mapping intact. (Academy required safe linking workflow.)
- **A07** Keep secrets/tokens server-side, bound outbound calls and sanitize errors; do not use provider tokens as forum session cookies. Prep: P13, P15. Evidence: Timeout/oversize scenarios fail safely; secret-marker scan is clean. (Academy operations detail.)
- **A08** Issue the existing forum session after identity resolution, enforce old-session invalidation and preserve all F01 guards and tests. Prep: P10, P15. Evidence: SSO login in browser B revokes browser A; base regression passes. (Source requirement.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- Two real SSO provider flows integrated into the forum
- Provider setup documentation, identity migration and callback/security regression suite

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **Two providers:** Google and GitHub first+repeat login → working registration and stable returning account.
- **Replay:** reuse completed callback → 400; no new session.
- **Cross browser:** callback state used without originating browser binding → 400.
- **Bad ID token:** wrong Google audience or nonce → rejected.
- **Collision:** provider email equals another local member → explicit proof required; no takeover.
- **Link conflict:** provider identity owned by another member → 409.
- **Session parity:** SSO login after local login in another browser → old session revoked.
- **Provider failure:** denial, timeout, malformed profile → safe error; base forum remains usable.

## Full guided chapters

[Open the HTML guide](../projects/f02-forum-authentication-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../authentication.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum/authentication). Source Go restrictions are replaced explicitly by this academy language choice.

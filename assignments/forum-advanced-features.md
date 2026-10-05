# Forum Advanced Features: Notifications, Activity and Editing

Optional · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 1 week after preparation

Integrate durable notifications, private personal activity and owner editing/deletion without duplicating events, leaking content or losing concurrent edits.

## Preparation gate

- F01: Forum: Full Stack Community
- P24: Owned Editing, Deletion and Stale Update Conflicts
- P25: Transactional Events and Notification Inbox
- P26: Personal Activity Queries and Index Observatory

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **V01** Notify post owners when other members comment or change a post reaction to like/dislike; do not notify self-actions or unchanged repeated reactions. Prep: P07, P25. Evidence: One correct recipient notification per meaningful change. (Source requirement.)
- **V02** Persist notification intent with the domain change, deliver idempotently and authorize recipient-only reads/read-marking. Prep: P25. Evidence: Rollback creates no event; dispatcher replay creates no duplicate. (Academy event integrity detail.)
- **V03** Provide durable offline notification inbox with read/unread state; normal refresh is the required transport and optional real-time delivery must have separate completed prep. Prep: P25. Evidence: Offline recipient sees committed notifications on next visit. (Academy delivery contract; source does not mandate a push protocol.)
- **V04** Activity page shows the logged-in user’s authored posts and where they currently liked/disliked posts or comments. Prep: P08, P26. Evidence: Cleared/switched reactions move to correct current sections. (Source requirement.)
- **V05** Activity page shows where and what the user commented, including comment text plus parent post context and a permitted link or neutral tombstone. Prep: P24, P26. Evidence: Comment fixture includes both content and parent context. (Source requirement.)
- **V06** Provide owner-only edit/remove sections for both posts and comments, with server ownership checks and CSRF protection. Prep: P10, P24. Evidence: Nonowner spoof fails; owner edits/removes both content types. (Source requirement.)
- **V07** Reject stale edits with 409 using expected versions; soft deletion applies the prepared child, reaction, asset and notification visibility policy. Prep: P24. Evidence: Two editors cannot silently overwrite; deleted parents do not leak. (Academy concurrency/lifecycle detail.)
- **V08** If moderation is installed, editing published content requeues it for review and withdraws public/asset visibility until approved; privileged overrides follow its matrix. Prep: P22, P23, P24. Evidence: Moderated edit becomes pending; no public leak. (Conditional integration requirement.)
- **V09** Use session-derived activity ownership, stable bounded pagination, correct SQL and justified indexes; retain all installed-feature tests. Prep: P11, P26. Evidence: Actor spoof fails; exact rows unchanged by index optimization. (Academy privacy/query detail.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- A notification inbox, personal activity sections and edit/delete pages
- Event, lifecycle and query-plan evidence plus regression tests

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **Comment/like/dislike:** B acts on A’s post → A receives correct durable notifications.
- **No-op/self:** repeat same reaction; A acts on own post → no extra notification.
- **Offline/retry:** A offline; dispatch same event twice → one notification available on return.
- **Privacy:** B requests A inbox/activity by supplied IDs → no A data; no state changes.
- **Activity context:** authored post/comment plus both target reactions → correct current sections and parent context.
- **Edit ownership/staleness:** nonowner edit then stale owner edit → 403 then 409; latest content intact.
- **Remove both types:** owner removes post and separate comment → public content hidden; lifecycle followed.
- **Installed moderation:** edit published post when F05 installed → pending; no public/asset leak.
- **Rollback/index regression:** failed domain write; optimize activity indexes → no phantom event; exact results retained.

## Full guided chapters

[Open the HTML guide](../projects/f06-forum-advanced-features-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../forum-advanced-features.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum/advanced-features). Source Go restrictions are replaced explicitly by this academy language choice.

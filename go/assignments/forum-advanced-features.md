# Forum Advanced Features: Notifications, Activity and Editing — Go

Optional · Go learning path · 4–5 students · 1 week after preparation

Integrate durable notifications, private personal activity and owner editing/deletion without duplicating events, leaking content or losing concurrent edits.

## Preparation gate

- [F01: Forum: Full Stack Community](../projects/f01-forum-guide/index.html)
- [P24: Owned Editing, Deletion and Stale Update Conflicts](../prep/p24-content-edit-lifecycle-guide/index.html)
- [P25: Transactional Events and Notification Inbox](../prep/p25-notification-event-inbox-guide/index.html)
- [P26: Personal Activity Queries and Index Observatory](../prep/p26-activity-query-observatory-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **V01** Notify post owners when other members comment or change a post reaction to like/dislike; do not notify self-actions or unchanged repeated reactions. Prep: P07, P25. Evidence: One correct recipient notification per meaningful change.
- **V02** Persist notification intent with the domain change, deliver idempotently and authorize recipient-only reads/read-marking. Prep: P25. Evidence: Rollback creates no event; dispatcher replay creates no duplicate.
- **V03** Provide durable offline notification inbox with read/unread state; normal refresh is the required transport and optional real-time delivery must have separate completed prep. Prep: P25. Evidence: Offline recipient sees committed notifications on next visit.
- **V04** Activity page shows the logged-in user’s authored posts and where they currently liked/disliked posts or comments. Prep: P08, P26. Evidence: Cleared/switched reactions move to correct current sections.
- **V05** Activity page shows where and what the user commented, including comment text plus parent post context and a permitted link or neutral tombstone. Prep: P24, P26. Evidence: Comment fixture includes both content and parent context.
- **V06** Provide owner-only edit/remove sections for both posts and comments, with server ownership checks and CSRF protection. Prep: P10, P24. Evidence: Nonowner spoof fails; owner edits/removes both content types.
- **V07** Reject stale edits with 409 using expected versions; soft deletion applies the prepared child, reaction, asset and notification visibility policy. Prep: P24. Evidence: Two editors cannot silently overwrite; deleted parents do not leak.
- **V08** If moderation is installed, editing published content requeues it for review and withdraws public/asset visibility until approved; privileged overrides follow its matrix. Prep: P22, P23, P24. Evidence: Moderated edit becomes pending; no public leak.
- **V09** Use session-derived activity ownership, stable bounded pagination, correct SQL and justified indexes; retain all installed-feature tests. Prep: P11, P26. Evidence: Actor spoof fails; exact rows unchanged by index optimization.

## Submission

- A notification inbox, personal activity sections and edit/delete pages
- Event, lifecycle and query-plan evidence plus regression tests

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f06-forum-advanced-features-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

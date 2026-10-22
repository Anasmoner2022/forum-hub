# Forum Moderation: Roles and Review Workflows — Go

Optional · Go learning path · 4–5 students · 1 week after preparation

Integrate guest/member/moderator/admin permissions, prepublication approval, role requests, reports, responses and category management into the forum.

## Preparation gate

- [F01: Forum: Full Stack Community](../projects/f01-forum-guide/index.html)
- [P22: Roles, Ownership and Permission Matrix](../prep/p22-role-policy-gateway-guide/index.html)
- [P23: Approval Queue, Role Requests, Reports and Categories](../prep/p23-moderation-workflow-desk-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **M01** Implement four access types: guests read published content; users create/comment/react; moderators review/delete/report posts; admins administer roles, reports, posts/comments and categories. Prep: P22. Evidence: Complete matrix tested through direct requests.
- **M02** Hold new posts pending until authorized approval; publish or reject through auditable transitions; authors see their own pending status and moderators cannot self-approve. Prep: P23. Evidence: Pending content is public only after approved transition.
- **M03** Moderators can delete existing posts and report posts to administrators with a stored reason. Prep: P22, P23. Evidence: Delete and report routes enforce moderator role.
- **M04** Users can request moderator status; admin can approve/reject the request and promote or demote users to/from moderator. Prep: P23. Evidence: Role request and audited decision persist.
- **M05** Admins receive reports, respond to them and resolve them; reporter can read the persisted response. Prep: P23. Evidence: Report→admin inbox→response→resolution demonstrated.
- **M06** Admins can delete both posts and comments. Deleted content is excluded from public views. Prep: P22, P23. Evidence: Direct comment deletion by admin succeeds; ordinary member denied.
- **M07** Admins can create and delete categories; an attached category deletion returns 409 under the prepared retention policy. Prep: P06, P23. Evidence: Unused category deleted; attached category remains intact.
- **M08** Bootstrap admin through the operator procedure; public registration cannot assign privileged roles; protect the final admin. Prep: P22. Evidence: Role spoof and final-admin race fail safely.
- **M09** Read current roles on every action, apply demotion immediately, guard all detail/feed/filter/count routes and installed image assets against pending/deleted leakage. Prep: P22, P23. Evidence: Old demoted cookie and guessed restricted IDs fail.
- **M10** Enforce expected-state/version transitions atomically with actor/reason/time audit; conflicting reviews return 409; retain F01 and installed-feature checks. Prep: P11, P23. Evidence: Simultaneous approve/reject has one winner and one conflict.

## Submission

- A moderation queue, member role-request form and admin desk
- Permission matrix, workflow migration and direct-request audit suite

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f05-forum-moderation-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

# Forum Moderation: Roles and Review Workflows

Optional · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 1 week after preparation

Integrate guest/member/moderator/admin permissions, prepublication approval, role requests, reports, responses and category management into the forum.

## Preparation gate

- F01: Forum: Full Stack Community
- P22: Roles, Ownership and Permission Matrix
- P23: Approval Queue, Role Requests, Reports and Categories

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **M01** Implement four access types: guests read published content; users create/comment/react; moderators review/delete/report posts; admins administer roles, reports, posts/comments and categories. Prep: P22. Evidence: Complete matrix tested through direct requests. (Source requirement.)
- **M02** Hold new posts pending until authorized approval; publish or reject through auditable transitions; authors see their own pending status and moderators cannot self-approve. Prep: P23. Evidence: Pending content is public only after approved transition. (Source requirement.)
- **M03** Moderators can delete existing posts and report posts to administrators with a stored reason. Prep: P22, P23. Evidence: Delete and report routes enforce moderator role. (Source requirement.)
- **M04** Users can request moderator status; admin can approve/reject the request and promote or demote users to/from moderator. Prep: P23. Evidence: Role request and audited decision persist. (Source requirement.)
- **M05** Admins receive reports, respond to them and resolve them; reporter can read the persisted response. Prep: P23. Evidence: Report→admin inbox→response→resolution demonstrated. (Source requirement.)
- **M06** Admins can delete both posts and comments. Deleted content is excluded from public views. Prep: P22, P23. Evidence: Direct comment deletion by admin succeeds; ordinary member denied. (Source requirement.)
- **M07** Admins can create and delete categories; an attached category deletion returns 409 under the prepared retention policy. Prep: P06, P23. Evidence: Unused category deleted; attached category remains intact. (Source requirement.)
- **M08** Bootstrap admin through the operator procedure; public registration cannot assign privileged roles; protect the final admin. Prep: P22. Evidence: Role spoof and final-admin race fail safely. (Academy privilege lifecycle detail.)
- **M09** Read current roles on every action, apply demotion immediately, guard all detail/feed/filter/count routes and installed image assets against pending/deleted leakage. Prep: P22, P23. Evidence: Old demoted cookie and guessed restricted IDs fail. (Academy complete visibility detail.)
- **M10** Enforce expected-state/version transitions atomically with actor/reason/time audit; conflicting reviews return 409; retain F01 and installed-feature checks. Prep: P11, P23. Evidence: Simultaneous approve/reject has one winner and one conflict. (Academy workflow integrity detail.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- A moderation queue, member role-request form and admin desk
- Permission matrix, workflow migration and direct-request audit suite

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **Guest/member separation:** direct privileged request from guest/member → 401/403; unchanged state.
- **Prepublication:** create then approve post → pending hidden; approved public.
- **Self approval:** moderator submits and reviews own post → 403.
- **Role lifecycle:** request, approve, demote, reuse old cookie → audit persists; demoted action denied.
- **Reports:** moderator report; admin response/resolution → reporter sees response.
- **Admin deletion:** admin deletes post and comment → both hidden publicly.
- **Categories:** admin creates, deletes unused, attempts attached delete → success, success, 409.
- **Conflict/leak audit:** concurrent review and guessed pending asset when installed → one review wins; restricted content denied.

## Full guided chapters

[Open the HTML guide](../projects/f05-forum-moderation-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../forum-moderation.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum/moderation). Source Go restrictions are replaced explicitly by this academy language choice.

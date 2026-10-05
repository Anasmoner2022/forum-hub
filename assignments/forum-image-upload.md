# Forum Image Upload: Verified Media in Posts

Optional · Node.js (JavaScript) **or** Go, chosen by the team · 4–5 students · 1 week after preparation

Allow registered members to create text posts with one verified image while preserving guest viewing, safe storage and container persistence.

## Preparation gate

- F01: Forum: Full Stack Community
- P16: Streaming Multipart Upload Gateway
- P17: Decoded Image Validation and Asset Lifecycle

Every applicable preparation verification case must pass before the assignment begins. Go teams satisfy equivalent behavioral gates with an approved compatible stack.

## Required behavior

- **I01** Registered users can create a post with text and one image; text-only posting still works and guests cannot upload. Prep: P16, P17. Evidence: Member uploads and guest rejection leave correct rows/files. (Source requirement.)
- **I02** Posts display their image to permitted readers, including guests when the parent post is public. Prep: P17. Evidence: Public post detail renders the verified image after restart. (Source requirement.)
- **I03** Support at least JPEG, PNG and GIF, including accepted animated GIFs; decode actual content rather than trusting extension or MIME. Prep: P17. Evidence: All formats pass; renamed text and corrupt images fail. (Source requirement.)
- **I04** Enforce 20,000,000 bytes per file; reject 20,000,001 with 413 and an image-too-big message. Whole request is at most 21,000,000 bytes; at most 8 fields, 1 file and 9 parts; each field at most 16 KiB. Native form repeats category up to five times. Prep: P16. Evidence: Byte-boundary and chunked-upload tests pass. (Source 20 mb made precise by academy decimal-MB policy.)
- **I05** Enforce P17 width, height, frame and total-pixel budgets with bounded decoding; oversized/corrupt content returns 422. Prep: P17. Evidence: Dimension and animation boundary tests pass. (Academy interpretation of source dimension caution.)
- **I06** Use generated private storage keys, safe verified serving headers, cleanup on every failure and reconciliation of crash leftovers; persist asset directories in Docker. Prep: P12, P16, P17. Evidence: Aborted upload and failed post linkage leave no public orphan. (Academy storage lifecycle detail.)
- **I07** Require session, origin and CSRF checks, preserve F01 atomic category creation, errors and tests; apply parent visibility if moderation is installed. Prep: P10, P16, P17. Evidence: Forged upload fails; guessed pending asset is denied in integrated mode. (Academy integration detail.)

## Backend and package policy

Choose Node.js or Go consistently for the forum and its extensions. Node.js worked guidance uses core HTTP/HTTPS, SQLite, crypto and testing APIs with prepared specialist packages. Go uses approved equivalents. No frontend framework, ORM or service replacing assessed domain logic. Read ../language-policy.html and ../runtime-policy.html.

## Submission

- Multipart post creation and protected asset-serving integration
- Required-format fixtures, byte/dimension limits and storage lifecycle evidence

Include a README, synthetic fixtures, tests with no required TODO/skips, migrations, decision record and peer reproduction evidence. Never submit secrets or real user credentials.

## Verification scenarios

- **JPEG/PNG/GIF:** upload one valid fixture each → post and image visible to guest if public.
- **Text only:** normal post without image → success with no asset.
- **File size:** 20,000,000 versus 20,000,001 bytes → intake accept versus 413; image validity separately checked.
- **Corrupt/spoof:** invalid bytes claimed as PNG → 422; cleanup.
- **Pixel/frame cap:** over stated dimension or frame budget → 422.
- **Abort/CSRF:** disconnect or bad CSRF field → no post/public asset; temp cleanup.
- **Storage failure:** final SQL linkage fails → compensated file; no partial post.
- **Visibility/restart:** replace container; request restricted asset in integrated mode → published retained; restricted denied.

## Full guided chapters

[Open the HTML guide](../projects/f03-forum-image-upload-guide/index.html). [Roadmap](../roadmap.html). [Coverage](../coverage.html).

Adapted from the supplied [original brief](../forum-image-upload.md) and [01-edu source subject](https://github.com/01-edu/public/tree/master/subjects/forum/image-upload). Source Go restrictions are replaced explicitly by this academy language choice.

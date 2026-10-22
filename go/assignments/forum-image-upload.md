# Forum Image Upload: Verified Media in Posts — Go

Optional · Go learning path · 4–5 students · 1 week after preparation

Allow registered members to create text posts with one verified image while preserving guest viewing, safe storage and container persistence.

## Preparation gate

- [F01: Forum: Full Stack Community](../projects/f01-forum-guide/index.html)
- [P16: Streaming Multipart Upload Gateway](../prep/p16-multipart-upload-gateway-guide/index.html)
- [P17: Decoded Image Validation and Asset Lifecycle](../prep/p17-image-asset-lifecycle-guide/index.html)

Complete every applicable Go preparation case before integration.

## Required behavior

- **I01** Registered users can create a post with text and one image; text-only posting still works and guests cannot upload. Prep: P16, P17. Evidence: Member uploads and guest rejection leave correct rows/files.
- **I02** Posts display their image to permitted readers, including guests when the parent post is public. Prep: P17. Evidence: Public post detail renders the verified image after restart.
- **I03** Support at least JPEG, PNG and GIF, including accepted animated GIFs; decode actual content rather than trusting extension or MIME. Prep: P17. Evidence: All formats pass; renamed text and corrupt images fail.
- **I04** Enforce 20,000,000 bytes per file; reject 20,000,001 with 413 and an image-too-big message. Whole request is at most 21,000,000 bytes; at most 8 fields, 1 file and 9 parts; each field at most 16 KiB. Native form repeats category up to five times. Prep: P16. Evidence: Byte-boundary and chunked-upload tests pass.
- **I05** Enforce P17 width, height, frame and total-pixel budgets with bounded decoding; oversized/corrupt content returns 422. Prep: P17. Evidence: Dimension and animation boundary tests pass.
- **I06** Use generated private storage keys, safe verified serving headers, cleanup on every failure and reconciliation of crash leftovers; persist asset directories in Docker. Prep: P12, P16, P17. Evidence: Aborted upload and failed post linkage leave no public orphan.
- **I07** Require session, origin and CSRF checks, preserve F01 atomic category creation, errors and tests; apply parent visibility if moderation is installed. Prep: P10, P16, P17. Evidence: Forged upload fails; guessed pending asset is denied in integrated mode.

## Submission

- Multipart post creation and protected asset-serving integration
- Required-format fixtures, byte/dimension limits and storage lifecycle evidence

Use reviewed Go modules, explicit SQL, synthetic fixtures, a reproducible Docker build and completed tests without required TODOs/skips.

[Full Go guide](../projects/f03-forum-image-upload-guide/index.html) · [Go roadmap](../roadmap.html) · [Go runtime policy](../runtime-policy.html)

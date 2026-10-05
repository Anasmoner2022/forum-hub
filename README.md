# Academy forum curriculum

Open [index.html](index.html) in a browser. Follow [roadmap.html](roadmap.html) and inspect [coverage.html](coverage.html).

- 26 preparation projects: 25 core + 1 optional live-database encryption lab.
- 6 full forum assignment guides: mandatory base, optional authentication, images, security, moderation and advanced features.
- Semantic HTML with a shared responsive dark theme and no browser JavaScript; unfinished downloadable student scaffolds.
- Worked language: Node.js. Team submission choice: Node.js or Go.
- Existing source Markdown briefs and the Learning Hub are left unchanged; use the adapted briefs in assignments/.

## Maintenance

Edit authored curriculum/*.mjs and run:

```text
npm run build
npm run check
```

The builder is dependency-free on Node 24.15+. It completes each whole project guide before moving to the next. Generated curriculum.json is the local inventory; learner completion is not inferred from content generation. Check reports verify curriculum structure, link integrity, coverage ordering and scaffold syntax, not completed student implementations.

Read the instructor release/preflight checklist before the first cohort. Exact dependency pins and time estimates require an academy pilot.

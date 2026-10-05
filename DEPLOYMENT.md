# Share the hub on Vercel

This repository builds a static curriculum website. Node.js runs during the build; students read the resulting HTML, CSS and starter downloads in their browsers.

## Deploy from GitHub

1. Commit and push the deployment changes to the GitHub repository connected to Vercel.
2. Open the Vercel project and use these settings:
   - Framework Preset: **Other**
   - Root Directory: the repository root (leave the field empty)
   - Build Command: **npm run build**
   - Output Directory: **public**
   - Node.js Version: **24.x**
3. Deploy the new commit. If Vercel already started a deployment after the push, use that deployment rather than redeploying the older failed commit.
4. Open the deployment URL and check a guide, its chapter links, the dark theme and a starter download. Share that URL with students.

The checked-in `vercel.json` selects Other, sets the build command and sets the output directory. No environment variables or runtime server are required for this hub.

## Verify locally

```text
npm run build
npm run check
npm run check:site
```

The build regenerates the existing guides, then creates `public/` with every HTML page and its linked assets, briefs and downloads. The published files retain the same directory layout and relative URLs. Build tools and authored curriculum modules stay outside `public/`.

`public/` is generated and ignored by Git. Vercel recreates it during each deployment. Continue opening the root `index.html` for the existing local workflow.

If the log says **No Output Directory named "public" found**, make sure the deployed commit contains the updated `package.json`, `tools/build-site.mjs` and `vercel.json`. The last build line should report **Published ... static files to public/**.

Official references: [Vercel build settings](https://vercel.com/docs/builds/configure-a-build) and [vercel.json configuration](https://vercel.com/docs/project-configuration/vercel-json).

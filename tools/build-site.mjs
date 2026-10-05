import {copyFile, lstat, mkdir, readFile, readdir, realpath, rm} from 'node:fs/promises';
import path from 'node:path';
import {fileURLToPath} from 'node:url';

const root = await realpath(path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..'));
const output = path.resolve(root, 'public');

// Only replace this project's generated output; never follow a linked directory.
if (path.dirname(output) !== root || path.basename(output) !== 'public') {
  throw new Error('The publish directory must be public inside this project.');
}
try {
  const info = await lstat(output);
  if (info.isSymbolicLink() || !info.isDirectory()) {
    throw new Error('Refusing to replace public: expected an ordinary directory.');
  }
} catch (error) {
  if (error.code !== 'ENOENT') throw error;
}

await import('./build-guides.mjs');
const inventory = JSON.parse(await readFile(path.join(root, 'curriculum.json'), 'utf8'));
const files = new Set(['curriculum.json']);
for (const entry of await readdir(root, {withFileTypes: true})) {
  if (entry.isFile() && entry.name.endsWith('.html')) files.add(entry.name);
}
for (const project of inventory.projects) {
  files.add(`${project.directory}/index.html`);
  for (const chapter of project.chapters) files.add(`${project.directory}/${chapter.file}`);
  for (const starter of project.starterFiles) files.add(`${project.directory}/starter/${starter}`);
}

// Follow local page links so briefs, CSS and downloads retain their existing URLs.
for (const relative of files) {
  const source = path.resolve(root, relative);
  if (!source.startsWith(root + path.sep)) throw new Error(`Publish path escapes the project: ${relative}`);
  if (!(await lstat(source)).isFile()) throw new Error(`Expected a regular publish file: ${relative}`);
  if (!relative.endsWith('.html')) continue;
  const html = await readFile(source, 'utf8');
  for (const match of html.matchAll(/(?:href|src)="([^"]+)"/g)) {
    const href = match[1].replaceAll('&amp;', '&');
    if (/^(?:[a-z][a-z0-9+.-]*:|\/\/|#)/i.test(href)) continue;
    const filename = decodeURIComponent(href.split(/[?#]/)[0]);
    if (!filename) continue;
    const target = path.resolve(path.dirname(source), filename);
    if (!target.startsWith(root + path.sep)) throw new Error(`Local link escapes the project: ${relative} -> ${href}`);
    files.add(path.relative(root, target).split(path.sep).join('/'));
  }
}

await rm(output, {recursive: true, force: true});
await mkdir(output, {recursive: true});
for (const relative of files) {
  const destination = path.join(output, relative);
  await mkdir(path.dirname(destination), {recursive: true});
  await copyFile(path.join(root, relative), destination);
}
console.log(`Published ${files.size} static files to public/ (${[...files].filter(file => file.endsWith('.html')).length} HTML pages).`);

import path from 'node:path';

const escape = value => String(value).replaceAll('&', '&amp;').replaceAll('<', '&lt;').replaceAll('>', '&gt;').replaceAll('"', '&quot;');
const hrefFrom = (directory, filename) => path.posix.relative(directory, filename);

export function renderPage(title, content, {directory = '', kind = 'document'} = {}) {
  const rootLink = filename => escape(hrefFrom(directory, filename));
  const sections = [];
  const body = content.replace(/<h2>(.*?)<\/h2>/g, (_, label) => {
    const id = `section-${sections.length + 1}`;
    sections.push({id, label});
    return `<h2 id="${id}">${label}</h2>`;
  });
  const outline = kind === 'document' && sections.length > 1
    ? `<aside class="page-outline" aria-label="On this page"><p class="outline-title">On this page</p><ol>${sections.map(({id, label}) => `<li><a href="#${id}">${label}</a></li>`).join('')}</ol></aside>`
    : '';
  const isHub = kind === 'hub';
  return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="color-scheme" content="dark">
<title>${escape(title)}</title>
<link rel="stylesheet" href="${rootLink('assets/hub.css')}">
</head>
<body class="${escape(kind)}">
<a class="skip-link" href="#main-content">Skip to content</a>
<header class="site-header">
  <div class="header-inner">
    <a class="brand" href="${rootLink('index.html')}" aria-label="Forum Academy — curriculum home"><span class="brand-mark" aria-hidden="true">F</span><span>Forum<span class="brand-subtitle">ACADEMY</span></span></a>
    <nav class="site-nav" aria-label="Curriculum navigation">
      <a href="${rootLink('index.html')}"${isHub ? ' aria-current="page"' : ''}>Overview</a>
      <a href="${rootLink('roadmap.html')}"${title === 'Forum curriculum roadmap' ? ' aria-current="page"' : ''}>Roadmap</a>
      <a href="${rootLink('coverage.html')}"${title === 'Forum requirement coverage' ? ' aria-current="page"' : ''}>Coverage</a>
      <a href="${rootLink('instructor-guide.html')}"${title === 'Instructor delivery and assessment' ? ' aria-current="page"' : ''}>For instructors</a>
    </nav>
    <span class="header-label">PROJECT TRACK</span>
  </div>
</header>
<main id="main-content" class="${outline ? 'reading-layout' : 'single-layout'}">
  <article class="${isHub ? 'hub-content' : 'reading-content'}">
    ${body}
  </article>
  ${outline}
</main>
<footer class="site-footer"><p>Learn it. Build it. Explain it.</p><a href="${rootLink('source-review.html')}">Sources &amp; academy adaptations</a></footer>
</body>
</html>
`;
}

export function renderHub({preps, capstones, phaseNames, stageGroups, byId, esc, projectLink, link, para}) {
  const tools = [
    ['roadmap.html', '01', 'Follow the roadmap', 'Learn each skill before the forum assignment that uses it.'],
    ['coverage.html', '02', 'Check the requirements', 'See how every requirement connects to preparation and evidence.'],
    ['language-policy.html', '03', 'Choose your backend', 'Understand the shared Node.js or Go submission contract.'],
    ['runtime-policy.html', '04', 'Prepare your workspace', 'Check runtimes, libraries and setup before you start.'],
    ['instructor-guide.html', '05', 'Teach the track', 'Read readiness gates, assessment and the cohort checklist.'],
    ['source-review.html', '06', 'Explore the sources', 'Review the original briefs and the academy adaptations.']
  ];
  return `<section class="hub-hero">
    <p class="eyebrow"><span class="status-dot" aria-hidden="true"></span> THE FULL STACK LEARNING PATH</p>
    <h1>Forum project track<span class="hero-accent">Prepare. Then build.</span></h1>
    <p class="hero-description">Build the skills first, then bring them together in a complete forum. Hands-on Node.js preparation, focused assignments, and a clear path from your first server to advanced features.</p>
    <div class="hero-actions"><a class="button button-primary" href="roadmap.html">Start with the roadmap <span aria-hidden="true">↗</span></a><a class="button button-secondary" href="#project-library">Browse the projects <span aria-hidden="true">↓</span></a></div>
    <div class="hero-note"><span>NODE.JS OR GO</span><span>4–5 STUDENTS PER TEAM</span><span>BUILD BEFORE YOU MOVE ON</span></div>
  </section>
  <div class="stat-grid" aria-label="Track at a glance">
    <div class="stat"><strong>26</strong><span>Preparation projects</span></div>
    <div class="stat"><strong>06</strong><span>Forum assignments</span></div>
    <div class="stat"><strong>02</strong><span>Backend choices</span></div>
    <div class="stat"><strong>01</strong><span>Guided roadmap</span></div>
  </div>
  <section class="hub-section">
    <div class="section-heading"><div><p class="eyebrow">YOUR STARTING POINT</p><h2>Find your next step</h2></div><p>Everything you need to navigate the track.</p></div>
    <div class="quick-grid">${tools.map(([href, number, title, description]) => `<a class="quick-card" href="${href}"><span class="card-number">${number}</span><h3>${title}<span aria-hidden="true">↗</span></h3><p>${description}</p></a>`).join('')}</div>
  </section>
  <section class="hub-section">
    <div class="section-heading"><div><p class="eyebrow">PUT YOUR SKILLS TO WORK</p><h2>Six forum assignments</h2></div><p>One mandatory foundation. Five optional extensions.</p></div>
    <div class="assignment-grid">${stageGroups.map(([phase, range, id]) => {
      const project = byId.get(id);
      return `<a class="assignment-card phase-${phase}" href="${esc(project.directory)}/index.html"><div class="assignment-top"><span class="project-id">${id}</span><span class="badge ${id === 'F01' ? 'badge-required' : ''}">${id === 'F01' ? 'Mandatory' : 'Optional'}</span></div><h3>${esc(phaseNames[phase])}</h3><p>${esc(project.goal)}</p><div class="assignment-meta"><span>${project.weeks} ${project.weeks === 1 ? 'week' : 'weeks'}</span><span>4–5 students</span></div><div class="assignment-bottom"><span>Prepare: ${esc(range)}</span><span aria-hidden="true">→</span></div></a>`;
    }).join('')}</div>
  </section>
  <section id="project-library" class="hub-section">
    <div class="section-heading"><div><p class="eyebrow">LEARN ONE THING. BUILD ONE THING.</p><h2>Preparation library</h2></div><p>Complete each group before its forum assignment.</p></div>
    <div class="library-grid">${Object.entries(phaseNames).map(([phase, title], index) => `<section class="phase-card phase-${phase}"><div class="phase-heading"><span class="phase-number">0${index + 1}</span><h3>${esc(title)}</h3><span class="phase-count">${preps.filter(p => p.phase === phase).length} projects</span></div><ol class="project-list">${preps.filter(p => p.phase === phase).map(p => `<li><a href="prep/${p.id.toLowerCase()}-${esc(p.slug)}-guide/index.html"><span class="project-id">${p.id}</span><span class="project-title">${esc(p.title)}</span><span class="project-time">${p.hours}h${p.bonus ? ' · optional' : ''}</span></a></li>`).join('')}</ol></section>`).join('')}</div>
  </section>
  <section class="hub-section info-grid">
    <div class="info-panel"><p class="eyebrow">MADE FOR BUILDING</p><h2>More than a reading list</h2>${para('Each guide includes exact reading, a focused refresher, a full specification, downloadable TODO scaffolds, guided build chapters, concrete verification, engineering questions and optional extensions. The student supplies the implementation or design.')}</div>
    <div class="info-panel"><p class="eyebrow">BEFORE YOU START</p><h2>Bring the basics</h2>${para('Start with basic JavaScript, HTML, terminal use and Git. Node.js and database concepts are taught here. The guided route uses Node.js; teams can choose Node.js or Go under the same behavioral contract. One preparation project is an optional encryption lab.')}</div>
  </section>
  <p class="storage-note">This is your standalone forum curriculum. All guides stay in this folder; the Learning Hub catalog and student progress are unchanged.</p>`;
}

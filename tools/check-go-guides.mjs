import {readFile, stat} from 'node:fs/promises';
import path from 'node:path';
import {tmpdir} from 'node:os';
import {spawnSync} from 'node:child_process';

export async function checkGoGuides(root, nodeInventory, check) {
 const inventory=JSON.parse(await readFile(path.join(root,'go/curriculum.json'),'utf8'));
 const ids=new Map(inventory.projects.map(p=>[p.id,p]));
 check(ids.size===32,'Go: expected 32 full guides');
 check(inventory.projects.filter(p=>p.kind==='prep').length===26,'Go: expected 26 preparation projects');
 check(inventory.projects.filter(p=>p.kind==='capstone').length===6,'Go: expected six assignments');
 check(JSON.stringify(inventory.requirements)===JSON.stringify(nodeInventory.requirements),'Go: behavior coverage must match the Node.js path');
 check(nodeInventory.tracks?.map(t=>t.id).join(',')==='nodejs,go','Expected Node.js and Go track inventory');
 const home=await readFile(path.join(root,'index.html'),'utf8');
 const goHome=await readFile(path.join(root,'go/index.html'),'utf8');
 for(const [language,text]of [['Node.js',home],['Go',goHome]]) {
  check(text.includes('Choose your backend language'),`${language}: missing backend selector`);
  check((text.match(/class="language-option"/g)||[]).length===2,`${language}: expected two backend choices`);
  const choices=text.match(/<nav class="language-options"[\s\S]*?<\/nav>/)?.[0]||'';
  check((choices.match(/aria-current="page"/g)||[]).length===1,`${language}: expected one selected backend`);
  const library=text.match(/<div class="library-grid">([\s\S]*?)<section class="hub-section info-grid">/)?.[1]||'';
  check((library.match(/class="project-title"/g)||[]).length===26,`${language}: expected 26 matching preparation links`);
  check(text.includes(`${language} preparation library`),`${language}: missing selected-language library heading`);
 }
 check(home.includes('href="go/index.html"'),'Node.js selector must open the Go path');
 check(goHome.includes('href="../index.html"'),'Go selector must open the Node.js path');
 check(goHome.includes('Go Language, Runtime and Module Lab')&&!goHome.includes('Node Runtime and Package Lab'),'Go library must use the Go foundation');
 check(home.includes('Node Runtime and Package Lab')&&!home.includes('Go Language, Runtime and Module Lab'),'Node.js library must use the Node.js foundation');

 const candidates=[process.env.GO_BINARY,'go',...(process.platform==='win32'?[path.join(process.env.ProgramFiles||'C:/Program Files','Go','bin','go.exe')]:[])].filter(Boolean);
 const go=candidates.find(exe=>spawnSync(exe,['version'],{encoding:'utf8',windowsHide:true}).status===0);
 check(Boolean(go),'Go compiler unavailable: install the cohort Go toolchain or set GO_BINARY to compile-check the starters');
 let compiled=0;
 const ordered=['P01','P02','P03','P04','P05','P06','P07','P08','P09','P10','P11','P12','F01','P13','P14','P15','F02','P16','P17','F03','P18','P19','P20','P21','F04','P22','P23','F05','P24','P25','P26','F06'];
 const done=new Set();
 for(const id of ordered) {
  const p=ids.get(id);
  check(Boolean(p),`Go: missing ${id}`);
  if(!p)continue;
  for(const dep of p.depends)check(done.has(dep),`Go: roadmap puts ${id} before ${dep}`);
  done.add(id);
 }
 for(const p of inventory.projects) {
  check(p.language==='go',`${p.id}: wrong Go language tag`);
  check(p.directory.startsWith('go/'),`${p.id}: Go directory must be on its own path`);
  const names=new Set(p.chapters.map(c=>c.file));
  check(names.size>=13,`Go ${p.id}: incomplete guide`);
  for(const name of ['overview.html','required-reading.html','go-refresher.html','spec.html','starter-structure.html','engineering-questions.html','practical-questions.html','extra-features.html','references.html'])check(names.has(name),`Go ${p.id}: missing ${name}`);
  check(names.has(p.design?'review-checklist.html':'testing.html'),`Go ${p.id}: missing verification`);
  check(p.steps.length>=3&&p.steps.every(s=>s.hints.length>=2),`Go ${p.id}: incomplete guided components`);
  check(p.cases.length>=6&&p.debate.length>=6&&p.practical.length>=5&&p.extras.length>=4,`Go ${p.id}: insufficient cases/questions/extras`);
  check(p.reads.some(([key])=>key.startsWith('go:'))||p.design,`Go ${p.id}: missing Go-specific required reading`);
  for(let i=0;i<p.steps.length;i++) {
   const text=await readFile(path.join(root,p.directory,`building-${String(i+1).padStart(2,'0')}.html`),'utf8');
   check(/<h2\b[^>]*>Self-check<\/h2>/.test(text),`Go ${p.id}: missing component self-check`);
  }
  for(const chapter of [{file:'index.html'},...p.chapters]) {
   const filename=path.join(root,p.directory,chapter.file),text=await readFile(filename,'utf8');
   check(!/node:|npm (?:ci|test|start|install)|eloquentjavascript\.net|nodejs\.org|Busboy|Sharp|DatabaseSync|@journeyapps|test\.todo/.test(text),`Go ${p.id}/${chapter.file}: Node.js-specific teaching leaked into Go guide`);
   for(const match of text.matchAll(/href="([^"]+\.html(?:#[^"]*)?)"/g)) {
    if(/^https?:/.test(match[1]))continue;
    const target=path.resolve(path.dirname(filename),match[1].split('#')[0]);
    check(target.startsWith(path.join(root,'go')+path.sep),`Go ${p.id}/${chapter.file}: chapter navigation leaves the Go path`);
   }
  }
  for(const file of p.starterFiles)check((await stat(path.join(root,p.directory,'starter',file))).isFile(),`Go ${p.id}: missing starter ${file}`);
  if(p.design)continue;
  check(p.starterFiles.includes('go.mod')&&p.starterFiles.includes('internal/lab/project.go'),`Go ${p.id}: missing Go scaffold`);
  check(!p.starterFiles.some(f=>f.endsWith('.js')||f==='package.json'),`Go ${p.id}: wrong-language scaffold`);
  const ignore=(await readFile(path.join(root,p.directory,'starter/.gitignore'),'utf8')).split(/\r?\n/);
  check(ignore.includes('/lab')&&!ignore.includes('lab'),`Go ${p.id}: executable ignore must not hide cmd/lab or internal/lab source directories`);
  const testText=await readFile(path.join(root,p.directory,'starter/internal/lab/project_test.go'),'utf8');
  check((testText.match(/t\.Fatal\("TODO:/g)||[]).length===p.cases.length,`Go ${p.id}: tests must start as explicit unfinished failures`);
  if(go) {
   const result=spawnSync(go,['test','-run','^$','./...'],{cwd:path.join(root,p.directory,'starter'),encoding:'utf8',windowsHide:true,timeout:120000,env:{...process.env,GOTOOLCHAIN:'local',GOCACHE:path.join(tmpdir(),'forum-hub-go-build-cache')}});
   compiled++;
   check(result.status===0,`Go ${p.id}: scaffold does not compile ${result.stderr||result.error||result.stdout}`);
  }
 }
 return {inventory,compiled};
}

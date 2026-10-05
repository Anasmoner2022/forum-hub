import {readFile,readdir,stat,writeFile} from 'node:fs/promises';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
import {spawnSync} from 'node:child_process';
const workspaceRoot=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'..');
const root=process.argv[2]?path.resolve(workspaceRoot,process.argv[2]):workspaceRoot;
const inventory=JSON.parse(await readFile(path.join(root,'curriculum.json'),'utf8'));
const errors=[];
const check=(ok,message)=>{if(!ok)errors.push(message);};
const ids=new Map(inventory.projects.map(p=>[p.id,p]));
check(ids.size===32,'Expected 32 unique guides');
check(inventory.projects.filter(p=>p.kind==='prep').length===26,'Expected 26 preparation guides');
check(inventory.projects.filter(p=>p.kind==='capstone').length===6,'Expected six final assignment guides');

async function walk(dir){const files=[];for(const e of await readdir(dir,{withFileTypes:true})){if(['node_modules','.git','.vercel','public'].includes(e.name))continue;const f=path.join(dir,e.name);if(e.isDirectory())files.push(...await walk(f));else files.push(f);}return files;}
const files=await walk(root),htmlFiles=files.filter(f=>f.endsWith('.html'));
const voidTags=new Set(['area','base','br','col','embed','hr','img','input','link','meta','param','source','track','wbr']);
let links=0,syntaxChecks=0;
for(const file of htmlFiles){
 const text=await readFile(file,'utf8'),label=path.relative(root,file);
 check(/^<!DOCTYPE html>/i.test(text),`${label}: missing doctype`);
 check(/<html lang="en">/.test(text),`${label}: missing language`);
 check(/<meta charset="UTF-8">/.test(text),`${label}: missing UTF-8 declaration`);
 check((text.match(/<h1>/g)||[]).length===1,`${label}: must have one h1`);
 check(!/<(?:script|style)\b|<[a-z][^>]*\s(?:style|on\w+)\s*=/i.test(text),`${label}: browser scripts or inline styling forbidden`);
 const stylesheets=[...text.matchAll(/<link rel="stylesheet" href="([^"]+)">/g)];
 check(stylesheets.length===1,`${label}: expected one shared dark-theme stylesheet`);
 for(const stylesheet of stylesheets)check(path.resolve(path.dirname(file),stylesheet[1])===path.join(root,'assets','hub.css'),`${label}: wrong theme path`);
 check(!/undefined|\[object Object\]|<p><\/p>/.test(text),`${label}: malformed generated content`);
 const stack=[];
 for(const match of text.matchAll(/<\/?([a-z][a-z0-9]*)(?:\s[^<>]*?)?\s*\/?>/gi)){
  const tag=match[1].toLowerCase();if(voidTags.has(tag))continue;
  if(match[0].startsWith('</')){const expected=stack.pop();check(expected===tag,`${label}: closing ${tag} but expected ${expected}`);}else stack.push(tag);
 }
 check(stack.length===0,`${label}: unclosed HTML tags ${stack.join(',')}`);
 const seenIds=new Set();for(const m of text.matchAll(/\bid="([^"]+)"/g)){check(!seenIds.has(m[1]),`${label}: duplicate id ${m[1]}`);seenIds.add(m[1]);}
 for(const m of text.matchAll(/href="([^"]+)"/g)){
  const href=m[1].replaceAll('&amp;','&');links++;
  check(!/^javascript:/i.test(href),`${label}: unsafe URL`);
  if(/^(https?:|mailto:)/.test(href))continue;
  const [target,fragment]=href.split('#');const resolved=path.resolve(path.dirname(file),decodeURIComponent(target||path.basename(file)));
  check(resolved===root||resolved.startsWith(root+path.sep),`${label}: local link escapes curriculum`);
  try{const s=await stat(resolved);check(s.isFile(),`${label}: link is not a file ${href}`);if(fragment){const dst=await readFile(resolved,'utf8');check(dst.includes(`id="${fragment}"`),`${label}: missing fragment ${href}`);}}catch{errors.push(`${label}: missing local target ${href}`);}
 }
}

function ancestors(id,visiting=new Set()){
 const p=ids.get(id);if(!p){errors.push(`Unknown prerequisite ${id}`);return new Set();}
 if(visiting.has(id)){errors.push(`Prerequisite cycle at ${id}`);return new Set();}
 const next=new Set(visiting);next.add(id);const result=new Set();
 for(const dep of p.depends){result.add(dep);for(const prior of ancestors(dep,next))result.add(prior);}return result;
}
for(const p of inventory.projects){
 ancestors(p.id);
 check(p.chapters.length>=13,`${p.id}: incomplete chapter set`);
 const chapterNames=new Set(p.chapters.map(c=>c.file));
 for(const name of ['overview.html','required-reading.html','node-refresher.html','spec.html','starter-structure.html','engineering-questions.html','practical-questions.html','extra-features.html','references.html'])check(chapterNames.has(name),`${p.id}: missing ${name}`);
 check(chapterNames.has(p.design?'review-checklist.html':'testing.html'),`${p.id}: missing verification chapter`);
 check(p.steps.length>=3,`${p.id}: too few component chapters`);
 check(p.cases.length>=6,`${p.id}: too few concrete cases`);
 check(p.debate.length>=6&&p.debate.length<=8,`${p.id}: expected 6–8 engineering questions`);
 check(p.practical.length>=5&&p.practical.length<=7,`${p.id}: expected 5–7 practical tasks`);
 check(p.extras.length>=4&&p.extras.length<=6,`${p.id}: expected 4–6 extensions`);
 for(let i=0;i<p.steps.length;i++){
  const f=path.join(root,p.directory,`building-${String(i+1).padStart(2,'0')}.html`);const content=await readFile(f,'utf8');
  check(/<h2\b[^>]*>Self-check<\/h2>/.test(content),`${p.id}: component ${i+1} lacks self-check`);
  check(p.steps[i].hints.length>=2,`${p.id}: component ${i+1} lacks grounded hints`);
 }
 for(const c of p.chapters){const content=await readFile(path.join(root,p.directory,c.file),'utf8');check(content.includes('← Index')&&content.includes('Previous:')||content.includes('← Index')&&c.file==='overview.html',`${p.id}/${c.file}: navigation incomplete`);}
 for(const filename of p.starterFiles){const f=path.join(root,p.directory,'starter',filename);const content=await readFile(f,'utf8');if(filename.endsWith('.js')){
   const result=spawnSync(process.execPath,['--check',f],{encoding:'utf8',windowsHide:true});syntaxChecks++;check(result.status===0,`${p.id}/${filename}: invalid scaffold syntax ${result.stderr}`);
  }else if(filename==='package.json'){try{JSON.parse(content);}catch{errors.push(`${p.id}: invalid starter package.json`);}}
 }
}

const reqIds=new Set();
for(const r of inventory.requirements){
 check(!reqIds.has(r.id),`Duplicate requirement ${r.id}`);reqIds.add(r.id);
 check(ids.has(r.capstone),`${r.id}: unknown final assignment`);
 check(r.prep.length>0,`${r.id}: no preparation mapping`);
 check(r.verify.length>15,`${r.id}: missing concrete evidence`);
 const available=ancestors(r.capstone);
 for(const id of r.prep){check(ids.get(id)?.kind==='prep',`${r.id}: invalid preparation ${id}`);if(!/Conditional|bonus/i.test(r.origin))check(available.has(id),`${r.id}: ${id} is not completed before ${r.capstone}`);}
}

// Check learning order independently of the file generation order.
const ordered=['P01','P02','P03','P04','P05','P06','P07','P08','P09','P10','P11','P12','F01','P13','P14','P15','F02','P16','P17','F03','P18','P19','P20','P21','F04','P22','P23','F05','P24','P25','P26','F06'];
const done=new Set();for(const id of ordered){for(const dep of ids.get(id).depends)check(done.has(dep),`Roadmap puts ${id} before ${dep}`);done.add(id);}
for(const p of inventory.projects.filter(p=>p.kind==='prep'&&!p.bonus))check(inventory.requirements.some(r=>r.prep.includes(p.id)||ancestors(r.capstone).has(p.id)),`${p.id}: unused prep`);

const report={checkedAt:new Date().toISOString(),passed:errors.length===0,guideCount:ids.size,htmlPages:htmlFiles.length,linksChecked:links,scaffoldSyntaxChecks:syntaxChecks,coverageRows:inventory.requirements.length,checks:['Semantic HTML, shared dark theme and complete chapter structure','Local links and balanced tags','Concrete cases, component self-checks and question counts','Prerequisite DAG and staged order','Requirement preparation mappings','Unfinished JavaScript scaffold syntax'],limitations:['No completed student solutions are supplied or certified','External URLs, live providers and native specialist packages require instructor preflight'],errors};
await writeFile(path.join(root,'validation-report.json'),JSON.stringify(report,null,2)+'\n');
console.log(JSON.stringify(report,null,2));
process.exitCode=errors.length?1:0;

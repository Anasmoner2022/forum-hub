// Local, read-only curriculum preview. Guide HTML contains no browser scripts.
import {createServer} from 'node:http';
import {readFile,stat} from 'node:fs/promises';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
const root=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'..');
const server=createServer(async(req,res)=>{
 try{
  if(req.method!=='GET'&&req.method!=='HEAD'){res.writeHead(405,{Allow:'GET, HEAD'});res.end();return;}
  const name=decodeURIComponent(new URL(req.url,'http://localhost').pathname);
  let filename=path.resolve(root,'.'+name);
  if(filename!==root&&!filename.startsWith(root+path.sep)){res.writeHead(403);res.end();return;}
  if((await stat(filename)).isDirectory())filename=path.join(filename,'index.html');
  const data=await readFile(filename);
  const types={'.html':'text/html; charset=utf-8','.css':'text/css; charset=utf-8','.json':'application/json; charset=utf-8','.csv':'text/csv; charset=utf-8'};
  res.writeHead(200,{'Content-Type':types[path.extname(filename)]||'text/plain; charset=utf-8','X-Content-Type-Options':'nosniff'});
  res.end(req.method==='HEAD'?undefined:data);
 }catch{res.writeHead(404,{'Content-Type':'text/plain'});res.end('Not found');}
});
server.listen(0,'127.0.0.1',()=>console.log(`Curriculum preview: http://127.0.0.1:${server.address().port}`));

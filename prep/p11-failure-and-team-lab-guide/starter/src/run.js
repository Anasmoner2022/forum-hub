import { runScenario } from './scenario.js';

try {
  await runScenario(process.argv.slice(2));
} catch (error) {
  console.error(error.message);
  process.exitCode = 1;
}

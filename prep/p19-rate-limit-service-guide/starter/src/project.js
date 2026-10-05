// Return {allowed, retryAfterSeconds} without extending blocked windows.
export async function allowAttempt(key, policy, now) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: allowAttempt');
}

// Remove stale keys while retaining live policy state.
export async function pruneLimiter(now) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: pruneLimiter');
}

// Enforce a bounded concurrent job count and safe overload behavior.
export async function runExpensiveJob(task) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: runExpensiveJob');
}

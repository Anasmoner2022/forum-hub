// Store a short-lived attempt and return the authorization URL.
export async function beginLoginAttempt(db, browserId, provider, clock) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: beginLoginAttempt');
}

// Claim exactly one valid attempt or reject.
export async function consumeLoginAttempt(db, browserId, state, clock) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: consumeLoginAttempt');
}

// Resolve token response with no automatic replay.
export async function exchangeCode(provider, code, verifier) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: exchangeCode');
}

// Return fresh token and cookie metadata after replacing old sessions.
export async function createSession(db, memberId, clock) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: createSession');
}

// Return trusted member context or null.
export async function authenticate(db, cookieHeader, clock) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: authenticate');
}

// Reject unsafe requests without the matching token and allowed origin.
export async function requireCsrf(request, session) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: requireCsrf');
}

// Revoke server state and expire the client cookie.
export async function logout(db, token) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: logout');
}

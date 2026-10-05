// Verify Google claims and return {provider,subject,email,emailVerified}.
export async function resolveGoogleIdentity(tokenResponse, attempt, config) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: resolveGoogleIdentity');
}

// Return stable GitHub ID with explicitly verified contact data.
export async function resolveGithubIdentity(accessToken, config) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: resolveGithubIdentity');
}

// Resolve a unique mapping; never auto-link by email.
export async function findOrCreateMember(db, identity, onboarding) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: findOrCreateMember');
}

// Attach only after recent reauthentication and conflict checks.
export async function linkIdentity(db, authenticatedMember, identity, linkIntent) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: linkIdentity');
}

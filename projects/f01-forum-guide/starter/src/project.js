// Map the documented routes to trusted guards and prepared services.
export async function dispatch(request, context) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: dispatch');
}

// Reuse P06 atomic creation with a session-derived author.
export async function createPost(context, input) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: createPost');
}

// Reuse prepared reply service with authorization.
export async function createComment(context, postId, body) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: createComment');
}

// Apply P07 contract to posts and comments.
export async function setForumReaction(context, target, value) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: setForumReaction');
}

// Return P08 public cards with current-user filter semantics.
export async function getForumFeed(context, filters) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: getForumFeed');
}

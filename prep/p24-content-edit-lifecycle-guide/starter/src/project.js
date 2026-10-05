// Apply validated owner edit conditionally.
export async function editOwnedPost(db, actor, postId, input, expectedVersion, policy) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: editOwnedPost');
}

// Apply owner edit or forbidden/conflict.
export async function editOwnedComment(db, actor, commentId, body, expectedVersion) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: editOwnedComment');
}

// Soft-delete and enforce child visibility consistently.
export async function deleteOwnedContent(db, actor, target, expectedVersion) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: deleteOwnedContent');
}

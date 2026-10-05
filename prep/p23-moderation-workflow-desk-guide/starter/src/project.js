// Apply one valid review and audit it atomically.
export async function reviewPost(db, actor, postId, decision, expectedVersion) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: reviewPost');
}

// Record decision and change role when approved.
export async function decideRoleRequest(db, admin, requestId, decision) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: decideRoleRequest');
}

// Persist admin response and resolution state.
export async function respondToReport(db, admin, reportId, response) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: respondToReport');
}

// Reject retained-post references with conflict.
export async function deleteCategory(db, admin, categoryId) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: deleteCategory');
}

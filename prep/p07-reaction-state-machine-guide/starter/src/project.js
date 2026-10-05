// Set the requested current value atomically; target={kind,id}.
export async function setReaction(db, memberId, target, value) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: setReaction');
}

// Remove the current value if present.
export async function clearReaction(db, memberId, target) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: clearReaction');
}

// Return {likes, dislikes} with zero defaults.
export async function reactionCounts(db, target) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: reactionCounts');
}

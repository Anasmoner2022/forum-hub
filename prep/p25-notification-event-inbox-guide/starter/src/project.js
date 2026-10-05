// Insert event inside the domain transaction with stable event ID.
export async function recordDomainEvent(db, event) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: recordDomainEvent');
}

// Deliver a bounded batch idempotently.
export async function dispatchEvents(db, limit = 100) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: dispatchEvents');
}

// Return only authorized recipient notifications.
export async function listInbox(db, recipientId, page) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: listInbox');
}

// Update only that recipient’s notification.
export async function markNotificationRead(db, recipientId, notificationId) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: markNotificationRead');
}

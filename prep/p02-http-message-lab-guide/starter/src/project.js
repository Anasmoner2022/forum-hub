// Return {status, headers, body}; consume request only for routes requiring a body.
export async function dispatch(request, context) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: dispatch');
}

// Resolve a Buffer within the byte limit or reject with a body error.
export async function readBoundedBody(request, limitBytes) {
  // TODO: implement this project's contract; see specification and guided chapters.
  throw new Error('TODO: readBoundedBody');
}

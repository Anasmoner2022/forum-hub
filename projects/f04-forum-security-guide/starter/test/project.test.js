import test from 'node:test';
import assert from 'node:assert/strict';
import * as project from '../src/project.js';

// TODO: replace every todo with an independently specified assertion.
// Each test owns its temp database, HTTP server, clock and cleanup.
void assert;
void project;
test.todo("HTTPS");
test.todo("Bad configuration");
test.todo("Login limits");
test.todo("Work concurrency");
test.todo("Session fixation");
test.todo("Revocation/expiry");
test.todo("CSRF/injection");
test.todo("Regression");

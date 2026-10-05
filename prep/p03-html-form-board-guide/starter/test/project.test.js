import test from 'node:test';
import assert from 'node:assert/strict';
import * as project from '../src/project.js';

// TODO: replace every todo with an independently specified assertion.
// Each test owns its temp database, HTTP server, clock and cleanup.
void assert;
void project;
test.todo("Creation");
test.todo("Blank title");
test.todo("Boundary title");
test.todo("Hostile body");
test.todo("Refresh");
test.todo("Missing item");

// Shared domain requirements stay in prep-base/prep-extensions. These are the Go
// tools, examples and component guidance used to implement those requirements.
export const goSources = {
  tour: ['A Tour of Go', 'https://go.dev/tour/welcome/1'],
  modules: ['Go: Create a module', 'https://go.dev/doc/tutorial/create-module'],
  errors: ['Go: Error handling and Go', 'https://go.dev/blog/error-handling-and-go'],
  os: ['Go: os', 'https://pkg.go.dev/os'],
  json: ['Go: encoding/json', 'https://pkg.go.dev/encoding/json'],
  http: ['Go: net/http', 'https://pkg.go.dev/net/http'],
  url: ['Go: net/url', 'https://pkg.go.dev/net/url'],
  io: ['Go: io', 'https://pkg.go.dev/io'],
  template: ['Go: html/template', 'https://pkg.go.dev/html/template'],
  sql: ['Go: Accessing databases', 'https://go.dev/doc/database/'],
  tx: ['Go: Executing transactions', 'https://go.dev/doc/database/execute-transactions'],
  sqlite: ['go-sqlite3: installation, DSN and database/sql driver', 'https://github.com/mattn/go-sqlite3'],
  test: ['Go: testing', 'https://pkg.go.dev/testing'],
  httptest: ['Go: net/http/httptest', 'https://pkg.go.dev/net/http/httptest'],
  race: ['Go: Data race detector', 'https://go.dev/doc/articles/race_detector'],
  context: ['Go: context', 'https://pkg.go.dev/context'],
  rand: ['Go: crypto/rand', 'https://pkg.go.dev/crypto/rand'],
  subtle: ['Go: crypto/subtle', 'https://pkg.go.dev/crypto/subtle'],
  scrypt: ['Go: golang.org/x/crypto/scrypt', 'https://pkg.go.dev/golang.org/x/crypto/scrypt'],
  oauth: ['Go: golang.org/x/oauth2', 'https://pkg.go.dev/golang.org/x/oauth2'],
  oidc: ['go-oidc: OpenID Connect verification', 'https://github.com/coreos/go-oidc'],
  multipart: ['Go: mime/multipart', 'https://pkg.go.dev/mime/multipart'],
  image: ['Go: image', 'https://pkg.go.dev/image'],
  gif: ['Go: image/gif', 'https://pkg.go.dev/image/gif'],
  exec: ['Go: os/exec', 'https://pkg.go.dev/os/exec'],
  tls: ['Go: crypto/tls', 'https://pkg.go.dev/crypto/tls'],
  sync: ['Go: sync', 'https://pkg.go.dev/sync'],
  cipher: ['go-sqlcipher/v4: encrypted SQLite driver', 'https://github.com/mutecomm/go-sqlcipher'],
  strings: ['Go: strings', 'https://pkg.go.dev/strings']
};

// Each reading names a concrete section, type or function, rather than an
// undirected link to an entire documentation site.
const read = (key, sections, why) => [`go:${key}`, sections, why];
const entries = {
  P01: {
    title: 'Go Language, Runtime and Module Lab', slug: 'go-runtime-lab', hours: 20,
    reads: [read('tour', 'Basics; Flow control; More types; Methods and interfaces; Concurrency', 'Learn Go syntax, pointers, slices, maps, methods, interfaces, errors, goroutines and channels before using them in later labs.'), read('modules', 'Create a module; Call your code from another module; Add a test; Compile and install', 'Separate the executable from an importable package.'), read('os', 'ReadFile; LookupEnv; Args; Exit', 'Read configuration and keep process exit at the CLI boundary.')],
    refresher: ['Go functions return errors as values. Read JSON into a typed struct, validate fields and return an error; main owns stderr and the process exit code. Complete the Tour exercises first, including interfaces and concurrency.', 'type Config struct {\n    AppName string `json:"appName"`\n    Port int `json:"port"`\n}\n// json.Unmarshal(data, &cfg) decodes structure; domain validation is separate.'],
    notes: ['Put CLI argument handling in cmd/lab and reusable code in internal/lab. Use os.LookupEnv to distinguish an absent PORT from an empty PORT.', 'Use os.ReadFile, encoding/json and strconv; handle returned errors explicitly. Distinguish bytes, UTF-8 text, rune counts, zero values and invalid configuration.', 'Teach go.mod, module paths, imports, gofmt, go test, go build, defer, errors.Is, contexts and synchronization. Commit go.sum once external dependencies are introduced.'],
    scope: ['Accept go run ./cmd/lab inspect fixtures/config.json; output JSON with appName, port and runtimeMajor="go1.26".', 'Validate appName as a nonempty string and port as an integer from 1 through 65535. PORT overrides the file only when present.', 'Missing files, invalid JSON, invalid fields and read failures return classified errors; main prints one short diagnostic to stderr and exits 1.', 'Use a Go module, typed configuration and an internal package. Commit go.mod and, when dependencies exist, go.sum. Keep files and secrets out of init functions.', 'Complete the named Tour sections and exercises for slices/maps, pointers, structs, methods, interfaces, errors, defer, goroutines, channels, select and mutexes. Explain when a sequential function is sufficient before adding concurrency.'],
    cases: [['Valid file', '{"appName":"inspector","port":3000}', 'JSON includes appName=inspector, port=3000 and runtimeMajor="go1.26"; exit 0'], ['Environment override', 'PORT=4000 with file port 3000', 'port=4000'], ['Malformed JSON', 'a file containing {', 'stderr identifies invalid JSON; exit 1'], ['Missing file', 'fixtures/missing.json', 'stderr identifies missing configuration; exit 1'], ['Bad port', 'PORT=abc or port=0', 'validation error; exit 1'], ['Read failure', 'inject a failing file reader', 'returned error reaches main once; exit 1']],
    steps: [
      {title:'Learn Go and separate the CLI', task:'Complete the named Tour exercises, create a module, and separate main from configuration validation.', hints:['Explain a pointer, slice, interface and returned error using your configuration types.', 'Keep os.Exit in main; demonstrate defer cleanup in the package.'], check:'Import the internal package in a test without printing, reading a file or starting a goroutine.'},
      {title:'Read and validate typed configuration', task:'Decode the JSON fixture, apply the optional environment override and return a validated report.', hints:['Distinguish JSON syntax, filesystem and field-validation failures.', 'strconv.Atoi succeeds for 0: explain why your port rule still rejects it.'], check:'The malformed-file and missing-file cases produce different diagnostics and the same nonzero exit status.'},
      {title:'Make modules and tests reproducible', task:'Format, compile and test the CLI, then have a teammate reproduce it.', hints:['A standard-library-only project does not need an invented dependency or go.sum.', 'Use a small channel exercise to explain cancellation and a mutex exercise to explain the race detector before later server labs.'], check:'A clean copy runs go test ./..., go build ./cmd/lab and the documented inspect command with the same fixture result.'}
    ]
  },
  P02: {
    reads:[read('http','Handler; ServeMux; Server; MaxBytesReader; Request.Context; MaxBytesError','Build explicit routes with bounded request work.'),read('io','ReadAll; LimitReader; ReadCloser','Read bytes without confusing a length header with a safety boundary.'),read('httptest','NewRequest; NewRecorder; NewServer','Send realistic requests and disconnects.')],
    refresher:['An http.Handler writes one response. A request context reports cancellation; MaxBytesReader bounds the bytes a handler can consume. Set server header/read/write/idle deadlines as a separate budget.', 'func health(w http.ResponseWriter, r *http.Request) {\n    w.Header().Set("Content-Type", "application/json")\n    // TODO: write the health response and handle the request contract.\n}'],
    notes:['Route by method and r.URL.Path; distinguish a known path with the wrong method from an unknown path.', 'Wrap r.Body with http.MaxBytesReader, close it, classify *http.MaxBytesError, and test chunked input using httptest.NewServer.', 'Handle ordinary errors explicitly; a recovery middleware handles unexpected panics only. Cancellation and timeout tests must leave /health usable.']
  },
  P03: {
    reads:[read('template','Security Model; ExecuteTemplate; contextual escaping','Render untrusted text safely without template.HTML casts.'),read('http','Request.ParseForm; PostForm; Redirect','Handle native forms and the 303 redirect flow.'),read('strings','TrimSpace','Normalize inputs before rune-count validation.')],
    refresher:['html/template escapes values for their HTML context. Pass typed page data to templates, keep submitted text as strings, and count Unicode code points rather than byte length.', 'type MessagePage struct {\n    Title string\n    Body string\n    Errors map[string]string\n}\n// {{.Body}} is escaped by html/template; do not cast user input to template.HTML.'],
    notes:['Use a bounded request body before ParseForm, and read r.PostForm so query parameters cannot override submitted fields.', 'Use utf8.RuneCountInString after trimming; render into a bytes.Buffer before sending headers so template failures can become a 500.', 'Use http.Redirect with http.StatusSeeOther after a successful mutation; reload the detail page to prove no duplicate write.']
  },
  P04: {
    reads:[], refresher:['An ERD describes entities, participation and cardinality independently of the programming language. Derive keys and junction tables before choosing Go storage structs.', 'Member (member_id PK) 1 ── authors ── N Topic (topic_id PK, author_id FK)\nTopic N ── TopicLabel ── N Label\n// Document the minimum label count as a transaction invariant.'],
    notes:['Keep the ERD independent of Go structs; a pointer field is not a database nullability policy.', 'Document the unique target/member reaction key and the one-active-session rule in the schema.', 'A database/sql adapter will implement these rules in P05–P07; distinguish a foreign key from a multi-row invariant.']
  },
  P05: {
    reads:[read('sql','Opening a database handle; Executing SQL statements; Querying for data','Understand sql.DB, rows and error ownership.'),read('sqlite','Installation; CGO_ENABLED; DSN foreign keys and busy timeout','Prepare the actual driver and C compiler before repository work.')],
    refresher:['sql.DB is a pool, not one connection. Foreign-key enforcement is per SQLite connection; configure it in the driver DSN and test a newly acquired connection. Teach CGO and the native compiler now.', 'db, err := sql.Open("sqlite3", dsn)\nif err != nil { return err }\ndefer db.Close()\n// PingContext confirms opening; sql.Open alone may not establish a connection.'],
    notes:['Use database/sql with github.com/mattn/go-sqlite3 and a reviewed version in go.mod/go.sum; prove CGO_ENABLED=1 and a compiler work on your host or Linux container.', 'Keep immutable numbered migration files; use a transaction and record the schema version only after every statement succeeds.', 'Configure _foreign_keys=on and a bounded _busy_timeout in the driver DSN; verify them on each opened connection and test restart persistence.']
  },
  P06: {
    reads:[read('tx','Best practices; sql.Tx; Commit; Rollback','Keep topic/category writes atomic.'),read('sql','Querying for multiple rows; Handling nullable column values; Avoiding SQL injection risk','Scan typed results and close cursors.')],
    refresher:['Start a transaction, defer rollback, and use only tx methods until commit. Calling db.Exec inside the transaction can escape its connection and break atomicity.', 'tx, err := db.BeginTx(ctx, nil)\nif err != nil { return err }\ndefer tx.Rollback()\n// TODO: perform related writes through tx, check errors, then commit.'],
    notes:['Define Go input and result structs from the specification; represent missing data explicitly rather than overloading zero IDs.', 'Use ExecContext/QueryContext placeholders; never interpolate submitted IDs or title text into SQL.', 'Close rows, check rows.Err, distinguish sql.ErrNoRows and verify a failed final label insert leaves no partial topic.']
  },
  P07: {
    reads:[read('tx','Executing transactions; transaction method boundaries','Keep reaction transitions and counts consistent.'),read('sync','Mutex; synchronization guarantees','Protect shared in-memory test coordination without replacing database constraints.')],
    refresher:['Concurrent handlers are goroutines. A mutex protects memory inside one process; a unique database key and a transaction protect persisted reaction state.', 'type Reaction int\nconst (\n    NoReaction Reaction = 0\n    Like Reaction = 1\n    Dislike Reaction = -1\n)\n// Enumerate all nine old/new transitions before implementing any of them.'],
    notes:['Use a typed reaction value and write all nine transitions before SQL.', 'Keep the target/member uniqueness constraint and execute changes on one sql.Tx; do not hold a transaction while calling external services.', 'Start competing requests with a barrier, wait for both, then read committed state; classify exhausted SQLite lock contention separately.']
  },
  P08: {
    reads:[read('sql','Querying data; Avoiding SQL injection risk','Build bounded, parameterized feed queries.'),read('url','Values; Query; Get','Parse submitted filters as untrusted strings.')],
    refresher:['Only fixed SQL fragments chosen by your program may change the query shape. Bind user values; derive the current actor from the session and keep pagination ordered by time plus ID.', 'type FeedFilter struct {\n    LabelID int64\n    Mine bool\n    Liked bool\n    Page int\n}\n// Parse and validate before building a query.'],
    notes:['Use r.URL.Query and strconv for bounded filters; document AND semantics in a typed filter.', 'Use EXISTS or preaggregated subqueries to avoid multiplying topic reactions by label rows.', 'Scan a bounded page into typed rows, check rows.Err and record query counts as well as EXPLAIN QUERY PLAN.']
  },
  P09: {
    reads:[read('scrypt','Key; CPU/memory cost parameters','Use a reviewed password derivation implementation.'),read('rand','Read','Generate independent salts.'),read('subtle','ConstantTimeCompare','Compare equal-length derived keys without a plain string equality check.')],
    refresher:['scrypt.Key returns a derived byte slice and an error. Bound record parameters before deriving; a goroutine alone is not a concurrency budget. Use a semaphore to cap simultaneous expensive work.', 'type HashRecord struct {\n    Version int\n    N, R, P int\n    Salt []byte\n    Digest []byte\n}\n// The record stores derivation parameters and a salt, never the original password.'],
    notes:['Normalize the email before the unique insert and handle a concurrent uniqueness violation; never trust a precheck as the only guard.', 'Use golang.org/x/crypto/scrypt, crypto/rand and crypto/subtle. Reject malformed parameters before allocating and derive with the trained work factors.', 'Use a bounded semaphore around hashing, measure peak memory, release it on every path and return one generic invalid-credentials response.']
  },
  P10: {
    reads:[read('http','Cookie; SetCookie; Request.Cookie','Separate opaque browser cookies from server state.'),read('rand','Read','Create unpredictable session and CSRF values.'),read('tx','Executing transactions','Enforce one active session per user atomically.')],
    refresher:['A session cookie carries an opaque random token, while expiry, actor and role live on the server. Use an injected Clock function in tests instead of sleeping.', 'type Clock func() time.Time\ntype Session struct {\n    UserID int64\n    ExpiresAt time.Time\n}\n// The browser cannot supply trusted UserID or role fields.'],
    notes:['Generate opaque tokens with crypto/rand, persist the prepared digest representation, and replace an earlier session atomically.', 'Use http.Cookie flags explicitly; test absolute and idle expiry using a fake clock and revoke by deleting the server record.', 'Implement CSRF and Origin checks as handlers/middleware before writes; compare tokens safely and reject cross-origin mutation attempts.']
  },
  P11: {
    reads:[read('test','T.Run; TempDir; Cleanup; Fatal; Error','Build table-driven assertions with isolated resources.'),read('httptest','NewServer; ResponseRecorder','Test HTTP behavior and cleanup.'),read('race','Usage; typical data races','Detect unsafe shared-memory access.')],
    refresher:['A Go test must fail when its invariant is violated. Use t.TempDir and t.Cleanup, join worker goroutines before the test returns, and check returned errors rather than discarding them.', 'func TestExample(t *testing.T) {\n    dir := t.TempDir()\n    _ = dir\n    // TODO: arrange independent input, exercise the boundary, assert the contract.\n    t.Fatal("TODO: replace this with real assertions")\n}'],
    notes:['Separate error classification from HTTP translation; unexpected panics produce a generic 500 and a redacted server diagnostic.', 'Use table-driven tests, httptest.NewServer, t.TempDir and t.Cleanup. Wait for concurrent work before asserting; use channel results rather than Fatal from a helper goroutine.', 'Run go test ./..., go vet ./... and go test -race ./... on a supported CGO platform; include migration conflicts and a deliberately injected failure.']
  },
  P12: {
    reads:[read('modules','Compile and install the application','Build an executable with explicit module inputs.'),read('sqlite','CGO_ENABLED and C compiler requirements','Keep builder and runtime native compatibility explicit.')],
    refresher:['A CGO SQLite binary may depend on the runtime C library. Build in a pinned Linux builder, run in a compatible runtime image and prove the final image starts as a non-root user.', '# Container build outline, to complete in this lab:\n# 1. Copy go.mod/go.sum and download reviewed modules.\n# 2. Build with CGO_ENABLED=1 in a compiler-equipped builder.\n# 3. Copy the executable into a compatible non-root runtime.'],
    notes:['Use a multi-stage Go Dockerfile, install the native compiler in the builder and record the libc/architecture requirements of the runtime.', 'Mount data and uploads as owned volumes; use exec-form startup and http.Server.Shutdown for graceful termination.', 'Use SQLite backup or VACUUM INTO under the prepared procedure rather than copying a live WAL database file; prove restored rows and constraints.']
  },
  P13: {
    reads:[read('http','Client; NewRequestWithContext; CheckRedirect; Client.Timeout','Bound outbound requests and destination trust.'),read('context','WithTimeout; cancellation','Propagate request deadlines.'),read('io','LimitReader','Bound provider response bytes independently of timeouts.')],
    refresher:['An http.Client should have an explicit timeout and redirect policy. Close response bodies, bound their bytes and validate the decoded provider shape before trusting it.', 'req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)\nif err != nil { return err }\n// TODO: apply the allowlisted destination, credential and response budget policy.'],
    notes:['Use httptest.NewServer for success, timeout, malformed JSON and redirect scenarios with synthetic credentials.', 'Configure http.Client timeout and CheckRedirect; use context cancellation, close resp.Body, and reject a response that exceeds the byte budget.', 'Keep secrets in server configuration; redact Authorization, tokens and code/state parameters from diagnostics.']
  },
  P14: {
    reads:[read('oauth','Config.AuthCodeURL; Config.Exchange; GenerateVerifier; S256ChallengeOption; VerifierOption','Prepare a standards-based code exchange and PKCE.'),read('rand','Read','Generate state and browser binding values.'),read('url','Values.Encode','Construct exact redirect parameters safely.')],
    refresher:['The browser carries an authorization code, not a trusted identity. State binds the attempt to the originating browser; a stored PKCE verifier is supplied only during the server token exchange.', 'verifier := oauth2.GenerateVerifier()\n// AuthCodeURL(..., oauth2.S256ChallengeOption(verifier))\n// Exchange(..., oauth2.VerifierOption(verifier))\n// TODO: bind state, browser and expiry before either call is usable.'],
    notes:['Draw the browser, forum and fake-provider sequence before wiring net/http handlers on separate ports.', 'Store a hashed browser-binding value, state, verifier and five-minute expiry; atomically consume the attempt using SQL before exchanging the code.', 'Use the reviewed golang.org/x/oauth2 options for S256; test provider errors, exact redirects and replays with the P13 client budgets.']
  },
  P15: {
    reads:[read('oidc','Provider discovery; IDTokenVerifier; Verify; Claims','Verify Google ID tokens with a maintained verifier and configured trust.'),read('oauth','Config.Exchange; Token; authenticated client','Exchange codes and call the GitHub identity endpoints.')],
    refresher:['An ID-token verifier checks signature and configured token policy; the application still binds the nonce to the consumed login attempt. GitHub identity comes from its authenticated API, not a decoded JWT.', 'type ProviderIdentity struct {\n    Issuer string\n    Subject string\n    Email string\n    EmailVerified bool\n}\n// Persist identity by issuer plus subject; linking is a separate authorized action.'],
    notes:['Register both provider test applications and demonstrate exact callback URIs before the forum assignment.', 'Use coreos/go-oidc/v3/oidc for Google discovery and verification; configure issuer, client ID and permitted signing algorithms, verify expiry, then compare the attempt nonce.', 'Use the bounded oauth2 client for GitHub /user and verified-email lookup; never auto-link accounts by matching email alone.']
  },
  P16: {
    reads:[read('multipart','Reader.NextPart; Part; streaming multipart parsing','Read one part at a time and enforce your own field/file counts.'),read('http','Request.MultipartReader; MaxBytesReader','Bound the whole request before multipart consumption.'),read('io','Copy; LimitReader','Bound individual fields and files.')],
    refresher:['MultipartReader streams parts; ParseMultipartForm may spill data to disk and is not a complete upload policy. Authenticate before consuming expensive data, enforce all counts and clean temporary files.', 'reader, err := r.MultipartReader()\nif err != nil { return err }\n// TODO: iterate NextPart, enforce counts/byte budgets and close every part.\n_ = reader'],
    notes:['Authenticate and check CSRF/Origin before accepting an upload; keep categories as repeated bounded form values.', 'Wrap the whole body in MaxBytesReader and enforce fields<=8, field bytes<=16 KiB, one file, parts<=9 and the distinct file/request limits.', 'Use os.CreateTemp outside the public directory, copy at most the file limit plus one byte, close handles and remove every rejected partial file.']
  },
  P17: {
    reads:[read('image','DecodeConfig; Decode; supported formats','Distinguish metadata from actual decoding.'),read('gif','DecodeAll; GIF frame fields','Validate every animation frame.'),read('exec','CommandContext; process lifetime and Wait','Isolate untrusted decoding work under resource limits.')],
    refresher:['JPEG, PNG and GIF decoders are registered through image/jpeg, image/png and image/gif. A small compressed file can expand into a large image; metadata checks and process resource limits are separate protections.', 'import (\n    "image"\n    _ "image/jpeg"\n    _ "image/png"\n    _ "image/gif"\n)\n// image.DecodeConfig is a preflight, not proof the entire image is valid.'],
    notes:['Use DecodeConfig for preflight, then fully decode JPEG/PNG or all GIF frames inside a constrained worker process. Goroutine cancellation cannot stop a decoder that ignores context.', 'Before decoding, inspect GIF descriptors with bounded parsing to reject excess frames/aggregate pixels; the worker must also have a timeout and a container/OS memory limit.', 'Use opaque server-generated names, an explicit MIME allowlist and parent visibility checks; compensate stored-file writes when the database transaction fails.']
  },
  P18: {
    reads:[read('tls','Config.MinVersion; LoadX509KeyPair','Configure a supported TLS floor and certificate/key loading.'),read('http','Server.ListenAndServeTLS; Shutdown','Serve HTTPS with explicit time budgets and graceful exit.')],
    refresher:['Keep the certificate and private key outside the image and repository. A SAN for localhost proves the development hostname; clients must trust the development certificate explicitly.', 'tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}\n// TODO: attach the configuration to the server and load mounted key material.\n_ = tlsConfig'],
    notes:['Generate the development SAN certificate with OpenSSL and document trust separately from certificate creation.', 'Use http.Server with TLSConfig.MinVersion=tls.VersionTLS12 and current safe defaults; never use InsecureSkipVerify to make a test pass.', 'Mount private keys read-only, verify permissions for the non-root process, and test Secure cookies and shutdown over HTTPS.']
  },
  P19: {
    reads:[read('sync','Mutex; map synchronization','Protect limiter state under concurrent handlers.'),read('context','Cancellation and deadlines','Bound the work admitted by the limiter.')],
    refresher:['net/http handles requests concurrently. Protect the limiter map with a mutex, inject a clock and separately cap the number of password derivations and live map entries.', 'type Bucket struct {\n    Used int\n    ResetAt time.Time\n}\n// One mutex-protected update must make the check and increment atomic.'],
    notes:['Use RemoteAddr as the default peer identity and only parse forwarding headers behind a documented trusted proxy.', 'Keep check/increment/reset atomic; use a fake clock and a start barrier to prove exactly the allowed concurrent requests succeed.', 'Bound map cardinality and eviction, use a semaphore for expensive hashes, and release permits on errors and cancellation.']
  },
  P20: {
    reads:[read('http','Cookie; SameSite; headers','Inspect what browsers actually receive.'),read('httptest','NewTLSServer; Client','Exercise trusted HTTPS clients without disabling verification.'),read('test','T.Run; Cleanup','Run deterministic attack fixtures.')],
    refresher:['A secure cookie flag is one layer: persisted revocation, CSRF, redirect validation and safe error responses still require server checks. Use NewTLSServer.Client for a correctly trusted development test client.', 'server := httptest.NewTLSServer(handler)\ndefer server.Close()\nclient := server.Client()\n// TODO: run the attack fixture using this configured client.\n_ = client'],
    notes:['Trace fixation, expiry, replay, role change, CSRF and open-redirect attempts before implementing each fixture.', 'Use a fake clock and persisted session rows, then combine concurrent requests with explicit synchronization and the race detector.', 'Inspect response headers and redacted diagnostics; test state changes and stored rows as well as status codes.']
  },
  P21: {
    reads:[read('cipher','Installation; import path; SQLCipher key usage','Review the actual encrypted engine and native build before selecting this optional lab.'),read('tx','Executing transactions','Retain database/sql transaction contracts under the encrypted driver.')],
    refresher:['Ordinary go-sqlite3 does not implement database encryption. Use a separately reviewed SQLCipher-enabled build, verify cipher_version and apply the key before any schema access on each connection.', '// Optional encrypted adapter contract:\n// 1. Load a key from a protected file.\n// 2. Open and key an SQLCipher connection before any read.\n// 3. Verify cipher_version and schema access, then expose repository operations.'],
    notes:['Build an isolated Linux CGO module with a reviewed SQLCipher driver and compiler. Do not import ordinary and encrypted drivers that register the same sqlite3 name into one executable.', 'Use a bounded single-connection sql.DB for the lab, set/key it before schema access and handle reconnection explicitly; transaction operations stay on one sql.Tx.', 'Use documented sqlcipher_export, verify schema metadata and foreign keys, then prove ordinary SQLite cannot read the encrypted target.', 'Rotate only disposable copies, verify old/new key behavior, and restore an encrypted snapshot with separate key custody.']
  },
  P22: {
    reads:[read('http','Handler; middleware composition','Apply authorization to direct requests as well as pages.'),read('tx','Executing transactions','Protect the final-admin invariant atomically.')],
    refresher:['Use typed actors/actions, a deny-by-default decision and the current persisted role. Request form fields do not supply trusted ownership or permissions.', 'type Actor struct {\n    ID int64\n    Role string\n}\n// TODO: define explicit actions and combine role, ownership and resource state.'],
    notes:['Translate every permission-matrix row into a table-driven Go test, including an unknown action.', 'Build trusted actor context from the session and current database role; never cache a browser-supplied role.', 'Put the final-admin count and mutation in one transaction, then run simultaneous demotion attempts.']
  },
  P23: {
    reads:[read('tx','Executing transactions; Commit and Rollback','Keep the moderation transition and audit write atomic.'),read('sql','ExecContext; RowsAffected; QueryContext','Detect stale decisions and build deterministic queues.')],
    refresher:['A conditional UPDATE with expected state/version lets one moderator win a race. Check RowsAffected and commit the audit record in the same transaction.', 'result, err := tx.ExecContext(ctx, statement, args...)\nif err != nil { return err }\n// TODO: distinguish one changed row from a stale or forbidden transition.\n_ = result'],
    notes:['Apply the same publication predicate to public pages, counts and asset handlers.', 'Represent role requests and reports as typed records with explicit status; retain an authorized report context when a post is deleted.', 'Use RowsAffected to detect a lost decision race, and roll back both state and audit when either write fails.']
  },
  P24: {
    reads:[read('sql','ExecContext; RowsAffected','Enforce expected-version ownership changes.'),read('tx','Executing transactions','Keep deletion and dependent state consistent.')],
    refresher:['Model a versioned edit as a conditional mutation on persisted author, ID and version. A stale save is a conflict, not a successful no-op.', 'type EditInput struct {\n    ExpectedVersion int64\n    Title string\n    Body string\n}\n// Actor comes from the session, not from this input.'],
    notes:['Use typed edit inputs and a trusted actor; check current ownership and the expected version in the write boundary.', 'Write the retained/deleted child visibility table before the transaction; keep public queries consistent with it.', 'Make the installed moderation policy explicit: the same validated edit remains published in base mode and becomes pending when moderation is installed.']
  },
  P25: {
    reads:[read('tx','Executing transactions','Persist domain changes and event intent together.'),read('sql','Querying multiple rows; executing changes','Drain bounded batches and check commit results.')],
    refresher:['Persist intent before delivery. A single-instance Go drain command can deliver durable inbox notifications without introducing a WebSocket server or an unbounded background goroutine.', 'type Event struct {\n    ID string\n    ActorID int64\n    RecipientID int64\n    Kind string\n}\n// The unique event/recipient key makes a delivery retry idempotent.'],
    notes:['Create events only for meaningful committed changes; reuse P07 transition fixtures to exclude repeated same-state and clear actions.', 'Keep notification insertion and delivered-marker update in one sql.Tx with a unique event-recipient constraint and a batch limit of 100.', 'Resolve current recipient and target visibility; use httptest to attempt another member’s list/read action by changing IDs.']
  },
  P26: {
    reads:[read('sql','Querying multiple rows; NULL handling; QueryContext','Scan private activity sections without N+1 lookups.'),read('url','Values.Get','Bound independent section pagination inputs.')],
    refresher:['The session actor defines whose activity is visible. Scan joined results into typed rows with nullable parent fields and apply current visibility before rendering a target link.', 'type ActivityRow struct {\n    Kind string\n    TargetID int64\n    Visible bool\n    Summary string\n}\n// A tombstone has no retained private summary or public target URL.'],
    notes:['Build expected activity rows from fixed fixtures before writing joined queries.', 'Keep separate bounded page parameters and pass the trusted actor ID to every query; use sql.NullString or an equivalent explicit nullable representation.', 'Capture EXPLAIN QUERY PLAN on identical 1000-record fixtures before/after indexes and instrument query counts.']
  }
};

const sharedLanguage = value => String(value)
  .replaceAll('A Node core HTTP server', 'A Go net/http server')
  .replaceAll('A native form and Busboy-based streaming intake service', 'A native form and Go streaming multipart intake service')
  .replaceAll('A TLS-enabled Node server', 'A TLS-enabled Go server')
  .replaceAll('a reproducible package lock', 'reproducible module metadata')
  .replaceAll('Await every HTTP operation', 'Wait for every HTTP operation to finish')
  .replaceAll('Buffer lengths', 'byte-slice lengths').replaceAll('a Buffer', 'a byte slice')
  .replaceAll('Use URL for', 'Use net/url for')
  .replaceAll('rejected repository promise', 'returned repository error')
  .replaceAll('awaited assertion observes failure', 'an explicit error assertion observes failure')
  .replaceAll('Remove an await and observe the test flaw.', 'Discard a returned error and observe the test flaw.')
  .replaceAll('Why can npm ci succeed locally but a container fail?', 'Why can a CGO build succeed locally but its Linux container fail?')
  .replaceAll('How does a callback driver change transaction sequencing?', 'How does a database/sql connection pool affect key initialization and transaction sequencing?')
  .replaceAll('open target with node:sqlite', 'open target with ordinary go-sqlite3');
const mapText = value => Array.isArray(value) ? value.map(mapText) : value && typeof value === 'object' ? Object.fromEntries(Object.entries(value).map(([k,v])=>[k,mapText(v)])) : typeof value === 'string' ? sharedLanguage(value) : value;
const unsupportedRead = key => key.startsWith('ej') || ['http','fs','sqlite','test','crypto','https','jose','busboy','sharp','sharpmeta','ciphernode'].includes(key);

export function prepareGo(project) {
  const p = mapText(project);
  const entry = entries[p.id];
  if (!entry) throw new Error(`Missing Go preparation: ${p.id}`);
  p.language = 'go';
  p.title = entry.title || p.title;
  p.slug = entry.slug || p.slug;
  p.hours = entry.hours || p.hours;
  p.reads = [...p.reads.filter(([key])=>!unsupportedRead(key)), ...entry.reads];
  p.refresher = entry.refresher;
  p.scope = entry.scope || p.scope;
  p.cases = entry.cases || p.cases;
  p.steps = entry.steps || p.steps.map((step,index)=>({...step, hints:[...step.hints,entry.notes[index]]}));
  p.goNotes = entry.notes;
  p.files = p.files.map(file=>file.replace(/^src\//,'internal/lab/').replace(/^test\//,'internal/lab/').replace(/\.test\.js$/,'_test.go').replace(/\.js$/,'.go').replaceAll('-','_'));
  p.api = p.api.map(([signature,description])=>[signature[0].toUpperCase()+signature.slice(1).split('(')[0],sharedLanguage(description)]);
  // Replace runtime-specific requirements explicitly; retain the shared domain contract.
  if(p.id==='P05') {
    p.scope[0]='Use database/sql with a reviewed github.com/mattn/go-sqlite3 version. Prepare CGO/compiler setup; configure and verify foreign-key enforcement on every connection.';
    p.scope[4]='Bound busy waits, close rows/connections, keep database files out of Git and prove data survives process restart. Demonstrate the same host/container driver behavior.';
  }
  if(p.id==='P07')p.scope[4]='Bound SQLite busy waits and map exhausted lock contention to a retriable error. Never perform external HTTP or other unbounded work while holding a database transaction.';
  if(p.id==='P09') {
    p.scope[2]='Use golang.org/x/crypto/scrypt with independent random salts and a versioned record. Training default: N=131072, r=8, p=1, 64-byte output, 16-byte salt. Bound simultaneous derivations to one initially and measure peak memory and latency before cohort launch.';
    p.scope[3]='Use crypto/subtle.ConstantTimeCompare on equal-length derived byte slices. Reject malformed/excessive record parameters before derivation; return generic invalid-credentials messages.';
  }
  if(p.id==='P11')p.scope[1]='Use testing and net/http/httptest, table-driven cases, t.TempDir/t.Cleanup, injected clocks and synchronized concurrent tests. Check all errors, join worker goroutines and run the race detector on a supported host.';
  if(p.id==='P12') {
    p.scope[0]='Use pinned, cohort-approved Go builder/runtime images. Build in Linux with the reviewed module versions, CGO_ENABLED=1 and a C compiler; document runtime libc and architecture compatibility.';
    p.scope[3]='Build .dockerignore excluding secrets, private keys, databases, local executables and unrelated academy guides. Never copy a Windows executable into a Linux runtime image.';
  }
  if(p.id==='P13')p.scope[0]='Call only a configured allowlisted provider endpoint with net/http.Client. Send server-side credentials only to that trusted origin and constrain redirects.';
  if(p.id==='P15')p.scope[1]='Google uses OpenID Connect: verify the ID token with a reviewed coreos/go-oidc/v3/oidc version and discovery JWKS. Configure permitted algorithms, issuer, audience/client ID and expiry, then compare the login-attempt nonce; key identity by issuer+subject.';
  if(p.id==='P16')p.scope[1]='Support multipart/form-data using Request.MultipartReader and mime/multipart. Enforce 8 fields, 16 KiB per field, one file, 9 total parts and 21,000,000 whole-request bytes. Reject unknown names and retain repeated category values before distinct-ID validation.';
  if(p.id==='P17')p.scope[0]='Accept fully decoded JPEG, PNG and GIF only. Use image/jpeg, image/png and image/gif. Preflight dimensions and GIF frame/pixel metadata with bounded parsing, then validate the full compressed stream in a worker with enforced memory/CPU/time budgets. DecodeConfig alone is insufficient.';
  if(p.id==='P18')p.scope[2]='Use net/http.Server with crypto/tls and a minimum of TLS 1.2, current safe runtime defaults, mounted key material and explicit timeouts. Never disable certificate verification globally.';
  if(p.id==='P21') {
    p.scope[0]='Optional only: ordinary SQLite cannot become encrypted through PRAGMA key. Prepare a separately reviewed github.com/mutecomm/go-sqlcipher/v4 build in an isolated Linux CGO module; the teacher verifies driver maintenance, engine version and supported platform before release.';
    p.scope[1]='Use database/sql with one bounded keyed connection for the lab and an explicit key-before-read procedure for every new connection. Do not register ordinary and SQLCipher sqlite3 drivers in the same executable. Teach CGO/native compiler and any selected build-mode dependencies before repository adaptation.';
    p.steps[1].title='Adapt database/sql without losing atomicity';
    p.steps[1].task='Retain P06/P07 contracts on the keyed database/sql connection and preserve one sql.Tx for each atomic operation.';
    p.steps[1].hints=['Check every returned error and commit result; do not interleave db calls with tx operations.', 'Configure key initialization before schema access and explicitly handle a replacement connection.', entry.notes[1]];
  }
  return p;
}

export function prepareGoCapstone(project) {
  const p = mapText(project);
  const prepByPhase = {base:['P05','P10','P11','P12'],authentication:['P13','P14','P15'],'image-upload':['P16','P17'],security:['P18','P19','P20'],moderation:['P22','P23'],advanced:['P24','P25','P26']};
  p.language = 'go';
  p.reads = [...p.reads.filter(([key])=>!unsupportedRead(key)), read('http','Handler; Server; Request.Context','Reuse the prepared Go HTTP boundary.'), read('test','T.Run; Cleanup','Repeat all accepted Go preparation cases.')];
  p.reads.push(...prepByPhase[p.phase].flatMap(id=>entries[id].reads));
  p.goNotes = prepByPhase[p.phase].map(id=>`${id}: reuse its tested Go component and preserve every required failure case.`);
  p.api = p.api.map(([signature,description])=>[signature[0].toUpperCase()+signature.slice(1).split('(')[0],sharedLanguage(description)]);
  p.files = p.files.map(file=>file.replace(/^src\//,'internal/lab/').replace(/^test\//,'internal/lab/').replace(/\.js$/,'.go'));
  p.steps = p.steps.map((step,index)=>({...step,hints:[...step.hints,p.goNotes[index%p.goNotes.length]]}));
  return p;
}

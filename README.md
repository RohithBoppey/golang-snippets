# golang-snippets

Small Go snippets, organised as reusable packages with tests, plus runnable demos.

## Layout

```
golang-snippets/
├── <topic>/            # reusable library package (e.g. scanning/)
│   ├── doc.go          # package notes
│   ├── <thing>.go      # functions (Capitalised = exported)
│   ├── <thing>_test.go # tests for <thing>.go
│   └── example_test.go # runnable examples, checked by `go test`
└── cmd/
    └── <demo>/main.go  # interactive programs (package main)
```

Current packages:

- `scanning` – reading input: `ReadLine`, `ScanWords`, `ScanInt`

## Commands

```sh
go test ./...                 # run all tests and examples
go test -v ./scanning         # verbose tests for one package
go doc ./scanning             # read package notes and docs
go run ./cmd/scan-multi       # run an interactive demo
go run ./cmd/scan-one
```

## Adding things

- **New topic:** create a folder `foo/` with `package foo`, a `.go` file and a `_test.go` file.
- **New function in a topic:** add `foo/bar.go` + `foo/bar_test.go`.
- **Reuse across topics:** `import "golang-snippets/foo"`.
- **New interactive demo:** create `cmd/<name>/main.go` that imports the package.

Rules: one folder = one package; `package main` only under `cmd/`; tests sit next to the code.

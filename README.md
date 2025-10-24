# projectstarter

Quick Go project scaffolding tool.

## Installation

```bash
go install github.com/alexshd/projectstarter/cmd/proj@latest
```

Or build from source:

```bash
git clone https://github.com/alexshd/projectstarter.git
cd projectstarter
go build -o proj ./cmd/proj
```

## Usage

```bash
# Create project with short name
proj start go myapp

# Create project with full module path
proj start go github.com/user/myapp
```

## What It Creates

- `cmd/projectname/main.go` - Working main with slog/tint setup
- `cmd/projectname/main_test.go` - Passing test
- `internal/` - Ready for your packages
- `go.mod` - Initialized with proper module path
- `README.md` - Basic documentation
- `LICENSE` - MIT license
- `.gitignore` - Go defaults

## Example

```bash
$ proj start go myapp
10:44:40 INF Creating Go project name=myapp
10:44:40 INF ✅ Project created successfully! name=myapp

$ cd myapp && go mod tidy && go run cmd/myapp/main.go
10:44:53 INF Starting myapp
Hello from myapp!

$ go test ./...
ok      myapp/cmd/myapp 0.002s
```

## License

MIT

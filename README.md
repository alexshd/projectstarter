# projectstarter

Quick Go project scaffolding tool.

> This is personal preferences and has no intention to be universal or create projects and setups that I don't use.

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
# Create Go project with short name
proj start go myapp

# Create Go project with full module path
proj start go github.com/user/myapp

# Create Vite + Elm + Tailwind project
proj start vite-elm myapp
```

## Project Types

### Go Project (`proj start go`)

Creates a Go project with:

- `cmd/projectname/main.go` - Working main with slog/tint setup
- `cmd/projectname/main_test.go` - Passing test
- `internal/` - Ready for your packages
- `go.mod` - Initialized with proper module path
- `README.md` - Basic documentation
- `LICENSE` - MIT license
- `.gitignore` - Go defaults

### Vite + Elm + Tailwind Project (`proj start vite-elm`)

Creates a modern Elm project with:

- Vite build setup with hot reload
- Elm with `vite-plugin-elm-watch` for instant feedback
- Tailwind CSS v4 with `@tailwindcss/vite` plugin
- `elm-tooling` for managing Elm tools (elm, elm-format, elm-json)
- Working counter example with Tailwind styling
- `package.json` with `dev`, `build`, `test` scripts
- `postinstall` hook to auto-install Elm tools

## Example

**Go Project:**

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

**Vite + Elm Project:**

```bash
$ proj start vite-elm my-elm-app
18:12:21 INF Creating Vite + Elm + Tailwind project name=my-elm-app
18:12:21 INF ✅ Project created successfully! name=my-elm-app

$ cd my-elm-app && npm install && npm run dev
# Elm tools installed via postinstall hook
# Dev server starts at http://localhost:5173
```

## License

MIT

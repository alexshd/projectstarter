package generator

import (
	"fmt"
	"time"
)

func (g *GoGenerator) mainGoTemplate(projectName string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func init() {
	// Initialize structured logging with colored output
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelInfo,
			TimeFormat: "15:04:05.0000",
			NoColor:    false,
			AddSource:  false,
		}),
	))
}

func main() {
	slog.Info("Starting %s")
	fmt.Println("Hello from %s!")
}
`, projectName, projectName)
}

func (g *GoGenerator) mainTestTemplate(projectName string) string {
	return `package main

import "testing"

func TestMain(t *testing.T) {
	// This test passes - you're ready to go!
	t.Log("Project initialized successfully")
}
`
}

func (g *GoGenerator) goModTemplate(modulePath string) string {
	return fmt.Sprintf(`module %s

go 1.21

require github.com/lmittmann/tint v1.1.2
`, modulePath)
}

func (g *GoGenerator) readmeTemplate(projectName, modulePath string) string {
	return fmt.Sprintf(`# %s

Created with projectstarter

## Installation

`+"```bash"+`
go mod tidy
`+"```"+`

## Usage

`+"```bash"+`
go run cmd/%s/main.go
`+"```"+`

## Testing

`+"```bash"+`
go test ./...
`+"```"+`

## License

MIT
`, projectName, projectName)
}

func (g *GoGenerator) licenseTemplate() string {
	year := time.Now().Year()
	return fmt.Sprintf(`MIT License

Copyright (c) %d

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`, year)
}

func (g *GoGenerator) gitignoreTemplate() string {
	return `# Binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output
*.out

# Go workspace file
go.work

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`
}

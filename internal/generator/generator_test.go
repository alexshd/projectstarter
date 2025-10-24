package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGoGenerator_Generate(t *testing.T) {
	t.Run("creates project successfully with short name", func(t *testing.T) {
		tmpDir := t.TempDir()
		oldDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}
		defer os.Chdir(oldDir)

		// Change to temp dir so project is created there
		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		gen := NewGoGenerator()
		if err := gen.Generate("myapp"); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		// Verify project directory exists
		if _, err := os.Stat("myapp"); os.IsNotExist(err) {
			t.Error("Project directory not created")
		}
	})

	t.Run("creates project successfully with module path", func(t *testing.T) {
		tmpDir := t.TempDir()
		oldDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}
		defer os.Chdir(oldDir)

		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		gen := NewGoGenerator()
		if err := gen.Generate("github.com/user/testapp"); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		// Should create directory with short name
		if _, err := os.Stat("testapp"); os.IsNotExist(err) {
			t.Error("Project directory not created")
		}
	})

	t.Run("fails when directory already exists", func(t *testing.T) {
		tmpDir := t.TempDir()
		oldDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}
		defer os.Chdir(oldDir)

		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		// Create directory first
		if err := os.MkdirAll("existing", 0o755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}

		gen := NewGoGenerator()
		err = gen.Generate("existing")
		if err == nil {
			t.Error("Expected error when directory exists, got nil")
		}
		if !strings.Contains(err.Error(), "already exists") {
			t.Errorf("Expected 'already exists' error, got: %v", err)
		}
	})

	t.Run("creates all required directories", func(t *testing.T) {
		tmpDir := t.TempDir()
		oldDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}
		defer os.Chdir(oldDir)

		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		gen := NewGoGenerator()
		if err := gen.Generate("test-dirs"); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		requiredDirs := []string{
			"test-dirs",
			filepath.Join("test-dirs", "cmd", "test-dirs"),
			filepath.Join("test-dirs", "internal"),
		}

		for _, dir := range requiredDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Errorf("Required directory not created: %s", dir)
			}
		}
	})

	t.Run("creates all required files", func(t *testing.T) {
		tmpDir := t.TempDir()
		oldDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}
		defer os.Chdir(oldDir)

		if err := os.Chdir(tmpDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		gen := NewGoGenerator()
		if err := gen.Generate("test-files"); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		requiredFiles := []string{
			"cmd/test-files/main.go",
			"cmd/test-files/main_test.go",
			"go.mod",
			"README.md",
			"LICENSE",
			".gitignore",
		}

		for _, file := range requiredFiles {
			fullPath := filepath.Join("test-files", file)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				t.Errorf("Required file not created: %s", file)
			}
		}
	})
}

func TestGoGenerator_ParseProjectName(t *testing.T) {
	gen := NewGoGenerator()

	tests := []struct {
		name               string
		input              string
		expectedModule     string
		expectedProjectDir string
	}{
		{
			name:               "short name",
			input:              "myapp",
			expectedModule:     "myapp",
			expectedProjectDir: "myapp",
		},
		{
			name:               "github module path",
			input:              "github.com/user/myapp",
			expectedModule:     "github.com/user/myapp",
			expectedProjectDir: "myapp",
		},
		{
			name:               "gitlab module path",
			input:              "gitlab.com/org/team/project",
			expectedModule:     "gitlab.com/org/team/project",
			expectedProjectDir: "project",
		},
		{
			name:               "custom domain module path",
			input:              "go.example.com/tools/app",
			expectedModule:     "go.example.com/tools/app",
			expectedProjectDir: "app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			modulePath, projectDir := gen.parseProjectName(tt.input)

			if modulePath != tt.expectedModule {
				t.Errorf("Expected module path %q, got %q", tt.expectedModule, modulePath)
			}
			if projectDir != tt.expectedProjectDir {
				t.Errorf("Expected project dir %q, got %q", tt.expectedProjectDir, projectDir)
			}
		})
	}
}

func TestGoGenerator_MainGo(t *testing.T) {
	gen := NewGoGenerator()
	content := gen.mainGoTemplate("testapp")

	t.Run("is valid package main", func(t *testing.T) {
		if !strings.Contains(content, "package main") {
			t.Error("main.go doesn't declare package main")
		}
	})

	t.Run("has main function", func(t *testing.T) {
		if !strings.Contains(content, "func main()") {
			t.Error("main.go doesn't have main function")
		}
	})

	t.Run("imports slog", func(t *testing.T) {
		if !strings.Contains(content, `"log/slog"`) {
			t.Error("main.go doesn't import log/slog")
		}
	})

	t.Run("imports tint", func(t *testing.T) {
		if !strings.Contains(content, `"github.com/lmittmann/tint"`) {
			t.Error("main.go doesn't import tint")
		}
	})

	t.Run("has init function for logging", func(t *testing.T) {
		if !strings.Contains(content, "func init()") {
			t.Error("main.go doesn't have init function")
		}
		if !strings.Contains(content, "slog.SetDefault") {
			t.Error("main.go doesn't configure default logger")
		}
	})

	t.Run("includes project name", func(t *testing.T) {
		if !strings.Contains(content, "testapp") {
			t.Error("main.go doesn't include project name")
		}
	})
}

func TestGoGenerator_MainTest(t *testing.T) {
	gen := NewGoGenerator()
	content := gen.mainTestTemplate("testapp")

	t.Run("is valid package main", func(t *testing.T) {
		if !strings.Contains(content, "package main") {
			t.Error("main_test.go doesn't declare package main")
		}
	})

	t.Run("imports testing", func(t *testing.T) {
		if !strings.Contains(content, `"testing"`) {
			t.Error("main_test.go doesn't import testing")
		}
	})

	t.Run("has test function", func(t *testing.T) {
		if !strings.Contains(content, "func TestMain(t *testing.T)") {
			t.Error("main_test.go doesn't have TestMain function")
		}
	})
}

func TestGoGenerator_GoMod(t *testing.T) {
	gen := NewGoGenerator()

	t.Run("with short name", func(t *testing.T) {
		content := gen.goModTemplate("myapp")

		if !strings.Contains(content, "module myapp") {
			t.Error("go.mod doesn't declare correct module")
		}
		if !strings.Contains(content, "go 1.21") {
			t.Error("go.mod doesn't specify Go version")
		}
		if !strings.Contains(content, "github.com/lmittmann/tint") {
			t.Error("go.mod doesn't require tint dependency")
		}
	})

	t.Run("with full module path", func(t *testing.T) {
		content := gen.goModTemplate("github.com/user/myapp")

		if !strings.Contains(content, "module github.com/user/myapp") {
			t.Error("go.mod doesn't declare correct module path")
		}
	})
}

func TestGoGenerator_Readme(t *testing.T) {
	gen := NewGoGenerator()
	content := gen.readmeTemplate("myapp", "github.com/user/myapp")

	t.Run("has project name as title", func(t *testing.T) {
		if !strings.Contains(content, "# myapp") {
			t.Error("README doesn't have project name as title")
		}
	})

	t.Run("has installation instructions", func(t *testing.T) {
		if !strings.Contains(content, "## Installation") {
			t.Error("README doesn't have Installation section")
		}
		if !strings.Contains(content, "go mod tidy") {
			t.Error("README doesn't mention go mod tidy")
		}
	})

	t.Run("has usage instructions", func(t *testing.T) {
		if !strings.Contains(content, "## Usage") {
			t.Error("README doesn't have Usage section")
		}
		if !strings.Contains(content, "go run cmd/myapp/main.go") {
			t.Error("README doesn't have correct run command")
		}
	})

	t.Run("has testing instructions", func(t *testing.T) {
		if !strings.Contains(content, "## Testing") {
			t.Error("README doesn't have Testing section")
		}
		if !strings.Contains(content, "go test ./...") {
			t.Error("README doesn't have test command")
		}
	})
}

func TestGoGenerator_License(t *testing.T) {
	gen := NewGoGenerator()
	content := gen.licenseTemplate()

	t.Run("is MIT license", func(t *testing.T) {
		if !strings.Contains(content, "MIT License") {
			t.Error("LICENSE doesn't identify as MIT")
		}
	})

	t.Run("has current year", func(t *testing.T) {
		currentYear := time.Now().Year()
		yearStr := fmt.Sprint(currentYear)
		if !strings.Contains(content, yearStr) {
			t.Errorf("LICENSE doesn't contain current year %d", currentYear)
		}
	})

	t.Run("has permission notice", func(t *testing.T) {
		if !strings.Contains(content, "Permission is hereby granted") {
			t.Error("LICENSE doesn't have permission notice")
		}
	})
}

func TestGoGenerator_Gitignore(t *testing.T) {
	gen := NewGoGenerator()
	content := gen.gitignoreTemplate()

	t.Run("ignores binaries", func(t *testing.T) {
		if !strings.Contains(content, "bin/") {
			t.Error(".gitignore doesn't ignore bin/")
		}
		if !strings.Contains(content, "*.exe") {
			t.Error(".gitignore doesn't ignore .exe files")
		}
	})

	t.Run("ignores test artifacts", func(t *testing.T) {
		if !strings.Contains(content, "*.test") {
			t.Error(".gitignore doesn't ignore test files")
		}
		if !strings.Contains(content, "*.out") {
			t.Error(".gitignore doesn't ignore output files")
		}
	})

	t.Run("ignores go.work", func(t *testing.T) {
		if !strings.Contains(content, "go.work") {
			t.Error(".gitignore doesn't ignore go.work")
		}
	})

	t.Run("ignores IDE files", func(t *testing.T) {
		if !strings.Contains(content, ".idea/") {
			t.Error(".gitignore doesn't ignore .idea/")
		}
		if !strings.Contains(content, ".vscode/") {
			t.Error(".gitignore doesn't ignore .vscode/")
		}
	})

	t.Run("ignores OS files", func(t *testing.T) {
		if !strings.Contains(content, ".DS_Store") {
			t.Error(".gitignore doesn't ignore .DS_Store")
		}
		if !strings.Contains(content, "Thumbs.db") {
			t.Error(".gitignore doesn't ignore Thumbs.db")
		}
	})
}

// TestGoGenerator_Integration performs an end-to-end test
func TestGoGenerator_Integration(t *testing.T) {
	tmpDir := t.TempDir()
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	gen := NewGoGenerator()
	if err := gen.Generate("integration-test"); err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	// Verify go.mod is valid by checking it's not empty and has module line
	goModPath := filepath.Join("integration-test", "go.mod")
	goModContent, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}
	if len(goModContent) == 0 {
		t.Error("go.mod is empty")
	}
	if !strings.Contains(string(goModContent), "module") {
		t.Error("go.mod doesn't contain module declaration")
	}

	// Verify main.go has valid Go syntax markers
	mainGoPath := filepath.Join("integration-test", "cmd", "integration-test", "main.go")
	mainGoContent, err := os.ReadFile(mainGoPath)
	if err != nil {
		t.Fatalf("Failed to read main.go: %v", err)
	}
	if !strings.Contains(string(mainGoContent), "package main") {
		t.Error("main.go doesn't have package declaration")
	}
	if !strings.Contains(string(mainGoContent), "func main()") {
		t.Error("main.go doesn't have main function")
	}

	// Verify all expected files exist and are non-empty
	requiredFiles := map[string]bool{
		"go.mod":                            true,
		"README.md":                         true,
		"LICENSE":                           true,
		".gitignore":                        true,
		"cmd/integration-test/main.go":      true,
		"cmd/integration-test/main_test.go": true,
	}

	for file := range requiredFiles {
		fullPath := filepath.Join("integration-test", file)
		info, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			t.Errorf("Required file doesn't exist: %s", file)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("File is empty: %s", file)
		}
	}
}

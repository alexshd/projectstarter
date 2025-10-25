package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type GoGenerator struct{}

func NewGoGenerator() *GoGenerator {
	return &GoGenerator{}
}

// Generate creates a new Go project with the given name
func (g *GoGenerator) Generate(projectName string) error {
	// Parse project name - could be "myapp" or "github.com/user/myapp"
	modulePath, projectDir := g.parseProjectName(projectName)

	// Check if directory already exists
	if _, err := os.Stat(projectDir); err == nil {
		return fmt.Errorf("directory '%s' already exists", projectDir)
	}

	// Create project structure
	if err := g.createStructure(projectDir, modulePath); err != nil {
		return err
	}

	return nil
}

// parseProjectName extracts module path and directory name
// "myapp" -> "myapp", "myapp"
// "github.com/user/myapp" -> "github.com/user/myapp", "myapp"
func (g *GoGenerator) parseProjectName(name string) (modulePath, projectDir string) {
	if strings.Contains(name, "/") {
		// Full module path provided
		modulePath = name
		parts := strings.Split(name, "/")
		projectDir = parts[len(parts)-1]
	} else {
		// Short name provided
		modulePath = name
		projectDir = name
	}
	return
}

// createStructure creates all project files and directories
func (g *GoGenerator) createStructure(projectDir, modulePath string) error {
	// Extract short name from path for use in templates
	shortName := filepath.Base(modulePath)

	// Create directories
	dirs := []string{
		projectDir,
		filepath.Join(projectDir, "cmd", shortName),
		filepath.Join(projectDir, "internal"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create files
	files := map[string]string{
		filepath.Join(projectDir, "cmd", shortName, "main.go"):      g.mainGoTemplate(shortName),
		filepath.Join(projectDir, "cmd", shortName, "main_test.go"): g.mainTestTemplate(shortName),
		filepath.Join(projectDir, "go.mod"):                         g.goModTemplate(modulePath),
		filepath.Join(projectDir, "README.md"):                      g.readmeTemplate(shortName, modulePath),
		filepath.Join(projectDir, "LICENSE"):                        g.licenseTemplate(),
		filepath.Join(projectDir, ".gitignore"):                     g.gitignoreTemplate(),
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", path, err)
		}
	}

	return nil
}

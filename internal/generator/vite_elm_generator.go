package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

type ViteElmGenerator struct{}

func NewViteElmGenerator() *ViteElmGenerator {
	return &ViteElmGenerator{}
}

// Generate creates a new Vite + Elm + Tailwind project
func (g *ViteElmGenerator) Generate(projectName string) error {
	// Check if directory already exists
	if _, err := os.Stat(projectName); err == nil {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	// Create project structure
	if err := g.createStructure(projectName); err != nil {
		return err
	}

	return nil
}

// createStructure creates all project files and directories
func (g *ViteElmGenerator) createStructure(projectName string) error {
	// Create directories
	dirs := []string{
		projectName,
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "public"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create files
	files := map[string]string{
		filepath.Join(projectName, "package.json"):     g.packageJSONTemplate(projectName),
		filepath.Join(projectName, "vite.config.js"):   g.viteConfigTemplate(),
		filepath.Join(projectName, "index.html"):       g.indexHTMLTemplate(projectName),
		filepath.Join(projectName, "src", "main.js"):   g.mainJsTemplate(),
		filepath.Join(projectName, "src", "style.css"): g.styleCSSTemplate(),
		filepath.Join(projectName, "src", "Main.elm"):  g.mainElmTemplate(),
		filepath.Join(projectName, "elm.json"):         g.elmJSONTemplate(),
		filepath.Join(projectName, "elm-tooling.json"): g.elmToolingJSONTemplate(),
		filepath.Join(projectName, ".gitignore"):       g.gitignoreTemplate(),
		filepath.Join(projectName, "README.md"):        g.readmeTemplate(projectName),
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", path, err)
		}
	}

	return nil
}

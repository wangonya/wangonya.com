package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/wangonya/wangonya.com/internal/generator"
	"github.com/wangonya/wangonya.com/internal/parser"
)

func processFiles(contentDir string, gen *generator.Generator) error {
	return filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Parse frontmatter and content
		meta, content, err := parser.ParseFrontmatter(content)
		if err != nil {
			return err
		}

		// Parse markdown
		parsed := parser.ParseMarkdown(content)

		// Generate HTML
		outputFile := filepath.Base(path[:len(path)-3]) + ".html"
		return gen.Generate(map[string]interface{}{
			"Title":   meta.Title,
			"Date":    meta.Date,
			"Tags":    meta.Tags,
			"Content": template.HTML(parsed.HTMLOutput),
		}, outputFile, strings.Contains(path, "/blog/"))
	})
}

func main() {
	// Define flags
	// contentDir := flag.String("content", "content", "Content directory path")
	// templateDir := flag.String("templates", "templates", "Templates directory path")
	// outputDir := flag.String("output", "output", "Output directory path")
	// watch := flag.Bool("watch", false, "Watch for file changes")
	// flag.Parse()

	// Initialize generator
	gen := generator.NewGenerator("templates", "public")

	// Process files
	err := processFiles("content", gen)
	if err != nil {
		log.Fatal(err)
	}

	// Serve
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Print("Listening on localhost:3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

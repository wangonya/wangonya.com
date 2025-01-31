package generator

import (
	"html/template"
	"os"
	"path/filepath"
)

type Generator struct {
	TemplateDir string
	OutputDir   string
}

func NewGenerator(templateDir, outputDir string) *Generator {
	return &Generator{
		TemplateDir: templateDir,
		OutputDir:   outputDir,
	}
}

func (g *Generator) Generate(data interface{}, outputFile string, blogPost bool) error {
	if err := os.MkdirAll(g.OutputDir+"/blog", 0755); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(filepath.Join(g.TemplateDir, "default.html"))
	if err != nil {
		return err
	}

	outputDir := g.OutputDir
	if blogPost {
		outputDir = outputDir + "/blog"
	}
	out, err := os.Create(filepath.Join(outputDir, outputFile))
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, data)
}

package render

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// PathRender renders a path template string (e.g. "cmd/{{.Project.K}}/main.go") with data.
func PathRender(pathTmpl string, data interface{}) string {
	t, err := template.New("path").Parse(pathTmpl)
	if err != nil {
		return pathTmpl
	}
	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return pathTmpl
	}
	return b.String()
}

// File renders a template file to dest with data.
func File(templatePath, destPath string, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("parse template %s: %w", templatePath, err)
	}
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}
	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", destPath, err)
	}
	defer f.Close()
	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	return nil
}

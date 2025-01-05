package template

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

var (
	//go:embed user/* layout/*
	templates      embed.FS
	layoutTemplate *template.Template
	once           sync.Once
)

func Render(w http.ResponseWriter, name string, data interface{}) error {
	once.Do(func() {
		// Parse only the layout user once
		layoutPath := filepath.ToSlash("layout/base.gohtml")
		content, err := templates.ReadFile(layoutPath)
		if err != nil {
			panic(fmt.Sprintf("failed to read layout user: %v", err))
		}

		tmpl, err := template.New(layoutPath).Parse(string(content))
		if err != nil {
			panic(fmt.Sprintf("failed to parse layout user: %v", err))
		}

		layoutTemplate = tmpl
	})

	tmpl, err := layoutTemplate.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone user: %w", err)
	}

	layoutPath := filepath.ToSlash(name)
	content, err := templates.ReadFile(layoutPath)
	if err != nil {
		return fmt.Errorf("failed to read layout user: %v", err)
	}

	tmpl, err = tmpl.New(layoutPath).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse layout user: %v", err)
	}

	// Set content type and write response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}

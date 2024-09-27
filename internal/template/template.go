package template

import (
	"embed"
	"html/template"
	"sync"
)

var (
	//go:embed "templates/*"
	templs   embed.FS
	renderer *Renderer
	onceDo   sync.Once
)

type Renderer struct {
	templ *template.Template
}




func NewRenderer() (*Renderer, error) {
	templ, err := template.ParseFS(templs, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	templ.

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

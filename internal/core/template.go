package core

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/starptech/go-web/internal/i18n"
)

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

type templateRenderer struct {
	templates map[string]*template.Template
}

// NewTemplateRenderer creates a new setup to render layout based go templates
func newTemplateRenderer(layoutsDir, templatesDir string) *templateRenderer {
	r := &templateRenderer{}
	r.templates = make(map[string]*template.Template)
	r.Load(layoutsDir, templatesDir)
	return r
}

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		c.Logger().Fatalf("the template %s does not exist", name)
		return fmt.Errorf("the template %s does not exist", name)
	}

	return tmpl.ExecuteTemplate(w, "base", data)
}

func (t *templateRenderer) Load(layoutsDir, templatesDir string) {
	layouts, err := filepath.Glob(layoutsDir)
	if err != nil {
		log.Fatal(err)
	}

	includes, err := filepath.Glob(templatesDir)

	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"Loc": i18n.Get,
	}

	mainTemplate := template.New("main")
	mainTemplate.Funcs(funcMap)

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range includes {
		fileName := filepath.Base(file)
		files := append(layouts, file)
		t.templates[fileName], err = mainTemplate.Clone()

		if err != nil {
			log.Fatal(err)
		}

		t.templates[fileName] = template.Must(t.templates[fileName].ParseFiles(files...))
	}
}

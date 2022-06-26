package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/captainmio/go-course/pkg/config"
	"github.com/captainmio/go-course/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultdata(data *models.TemplateData) *models.TemplateData {
	return data
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, gohtml string, data *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[gohtml]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	data = addDefaultdata(data)

	_ = t.Execute(buf, data) // pasing the data to the template

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("error writing template to browser:", err)
	}
}

// CreateTemplateCache create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		fmt.Println("error getting layout:", err)
		return templates, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			fmt.Println("error getting layout:", err)
			return templates, err
		}

		checkMatch, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			fmt.Println("error no match:", err)
			return templates, err
		}

		if len(checkMatch) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
		}

		templates[name] = templateSet
	}

	return templates, nil
}

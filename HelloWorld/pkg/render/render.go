package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, gohtml string) {

	_, err := RenderLayout(w)

	if err != nil {
		fmt.Println("error templateCache:", err)
		return
	}

	parseTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err = parseTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderLayout(w http.ResponseWriter) (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		fmt.Println("error getting layout:", err)
		return templates, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Current page", page)
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

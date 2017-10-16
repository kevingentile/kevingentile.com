package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates = ParseTemplates()

//ParseTemplates walks the template directory, parses all html templates, and returns a template pointer.
func ParseTemplates() *template.Template {
	tmp := template.New("")
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "template.html") {
			_, err = tmp.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return tmp
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = tmpl + ".template.html"                 //add html tag
	err := templates.ExecuteTemplate(w, tmpl, nil) //attempt to render the template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //Template not found/ not rendered
	}
}

package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"index.html", "left-sidebar.html", "right-sidebar.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = tmpl + ".html"                          //add html tag
	err := templates.ExecuteTemplate(w, tmpl, nil) //attempt to render the template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //Template not found/ not rendered
	}
}

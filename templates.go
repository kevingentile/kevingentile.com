package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"index.template.html", "left-sidebar.template.html", "banner.template.html", "carousel.template.html",
	"contact.template.html", "features.template.html", "footer.template.html", "photos.template.html",
	"header.template.html", "main.template.html", "scripts.template.html", "nav.template.html",
	"tweets.template.html", "posts.template.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = tmpl + ".template.html"                 //add html tag
	err := templates.ExecuteTemplate(w, tmpl, nil) //attempt to render the template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //Template not found/ not rendered
	}
}

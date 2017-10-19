package main

import (
	"github.com/LaughingCabbage/kevingentile.com/web"
	"net/http"
)

//makeHandler is a wrapper for all handler functions
func makeHandler(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handle(w, r)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	carousel := web.Carousel{
		Articles: []web.ReelArticle{
			{Title: "TEST ARTICLE", ImageLink: "#", Body: "test body"},
		},
	}
	tmpl := "index.template.html"                       //add html tag
	err := templates.ExecuteTemplate(w, tmpl, carousel) //attempt to render the template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //Template not found/ not rendered
	}
	//renderTemplate(w, "index")
}

func LeftHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "left-sidebar")
}

func RightHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "right-sidebar")
}

func NosideHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "no-sidebar")
}

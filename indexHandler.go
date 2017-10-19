package main

import (
	"net/http"

	"github.com/LaughingCabbage/kevingentile.com/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	carousel := web.Carousel{
		Articles: []web.ReelArticle{
			{Title: "TEST ARTICLE", ImageLink: "#", Body: "test body"},
		},
	}
	/*
		tmpl := "index.template.html"                       //add html tag
		err := templates.ExecuteTemplate(w, tmpl, carousel) //attempt to render the template
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //Template not found/ not rendered
		}
	*/
	renderTemplate(w, "index", Data{
		"carousel": carousel,
	})
}

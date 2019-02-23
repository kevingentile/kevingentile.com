package main

import (
	"net/http"

	"github.com/LaughingCabbage/kevingentile.com/web"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	//type Page struct {
	//	Title string
	//	Class string
	//}
	page := web.Page{
		Title: "Contact",
		Class: "no-sidebar",
	}
	/*
		//type Heading struct {
		//	Header string
		//	Body   string
		//}
		header := web.PageHeading{
			Header: "Contacts",
			Body:   "Email: gentile_kevin94@hotmail.com",
			Button: false,
		}
	*/

	main := web.PageMain{
		Heading: "Contact",
	}

	//type PageArticle struct {
	//	Heading    string
	//	Paragraphs []string
	//}
	articles := []web.PageArticle{
		{Heading: "Email",
			Paragraphs: []web.Paragraph{"gentile_kevin94@hotmail.com"},
		},
	}

	data := web.Data{"page": page, "main": main, "articles": articles, "nav": Nav, "footer": Footer, "scripts": ScriptsCommon}
	renderTemplate(w, "contact", data)
}

package main

import (
	"net/http"

	"github.com/LaughingCabbage/kevingentile.com/web"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {

	//type Page struct {
	//	Title string
	//	Class string
	//}
	page := web.Page{
		Title: "Links",
		Class: "no-sidebar",
	}

	main := web.PageMain{
		Heading: "Links",
		Body:    "A collection of links I find useful",
	}

	links := web.LinkList{
		Links: []web.PageLink{
			{Title: "Go Programming Language",
				Link:        "https://golang.org/",
				Description: "The new era of cross-platorm development via Google.",
			},
			{Title: "The Twelve-Factor App",
				Link:        "https://12factor.net/",
				Description: "The architecture for maintainable production applications",
			},
			{Title: "Git Branching Model",
				Link:        "http://nvie.com/posts/a-successful-git-branching-model/",
				Description: "A model for using git in development written by Vincent Driessen."},
			{Title: "High Performance Browser Networking",
				Link:        "https://hpbn.co/",
				Description: "The networking stack for developers written by Ilya Grigorik.",
			},
			{
				Title:       "Handling 1 million requests per minute with golang",
				Link:        "https://medium.com/smsjunk/handling-1-million-requests-per-minute-with-golang-f70ac505fcaa",
				Description: "An insight into to the power of Go for large scale development written by Marcio Castilho.",
			},
			{
				Title:       "Python Data Science Handbook",
				Link:        "https://jakevdp.github.io/PythonDataScienceHandbook/",
				Description: "Working with data in Python",
			},
		},
	}

	data := web.Data{"nav": Nav, "footer": Footer, "scripts": ScriptsCommon, "page": page, "main": main, "links": links}
	renderTemplate(w, "links", data)
}

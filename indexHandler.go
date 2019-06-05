package main

import (
	"net/http"

	"github.com/kevingentile/kevingentile.com/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//type Page struct {
	//	Title string
	//	Class string
	//}
	page := web.Page{
		Title: "Portfolio Home",
		Class: "homepage",
	}

	//type Heading struct {
	//	Header string
	//	Body   string
	//}
	heading := web.PageHeading{
		Header: ":)",
		Body:   "Welcome!",
		Button: true,
	}
	banner := web.PageHeading{
		Header: "Welcome to my portfolio!",
		Body:   "Here lies the fruits of my labor...",
	}

	//type ReelArticle struct {
	//	Title     string
	//	Body      string
	//	ImageLink string
	//	PicPath   string
	//}
	carousel := web.Carousel{
		Articles: []web.ReelArticle{
			{Title: "kevingentile.com",
				Body:      "This website! Built using Go and the HTML5UP site template.",
				ImageLink: "https://github.com/kevingentile/kevingentile.com",
				PicPath:   "../images/pic01.jpg",
			},
			{Title: "goLinks",
				Body:      "A blockchain library and data integrity tool built in Go",
				ImageLink: "https://github.com/kevingentile/goLinks",
				PicPath:   "../images/pic04.jpg",
			},
			{Title: "isoChar",
				Body:      "Image processing to for text tracing.",
				ImageLink: "https://github.com/kevingentile/isoChar",
				PicPath:   "../images/pic02.jpg",
			},
			{Title: "Arcade Game",
				Body:      "A game built with the SFML framework in C++",
				ImageLink: "https://github.com/kevingentile/Asteroid",
				PicPath:   "../images/pic03.jpg",
			},
		},
	}
	data := web.Data{"page": page, "heading": heading, "carousel": carousel, "nav": Nav, "footer": Footer, "banner": banner, "scripts": ScriptsCommon}
	renderTemplate(w, "index", data)

}

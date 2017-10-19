package main

import (
	"net/http"

	"github.com/LaughingCabbage/kevingentile.com/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//type Heading struct {
	//	Header string
	//	Body   string
	//}
	heading := web.Heading{
		Header: ":)",
		Body:   "Welcome!",
	}
	banner := web.Heading{
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
				ImageLink: "https://github.com/LaughingCabbage/kevingentile.com",
				PicPath:   "../images/pic01.jpg",
			},
			{Title: "Image Text Processing",
				Body:      "Find and draw a box around letters in an image.",
				ImageLink: "https://github.com/LaughingCabbage/isoChar",
				PicPath:   "../images/pic02.jpg",
			},
			{Title: "Arcade Game",
				Body:      "A game built with the SFML framework in C++",
				ImageLink: "https://github.com/LaughingCabbage/Asteroid",
				PicPath:   "../images/pic03.jpg",
			},
			{Title: "Blockchain",
				Body:      "Simple blockchain built in Go to better understand cryptocurrency",
				ImageLink: "https://github.com/LaughingCabbage/goLinks",
				PicPath:   "../images/pic04.jpg",
			},
		},
	}
	data := Data{"heading": heading, "carousel": carousel, "nav": nav, "footer": footer, "banner": banner, "scripts": scriptsCommon}
	renderTemplate(w, "index", data)
}

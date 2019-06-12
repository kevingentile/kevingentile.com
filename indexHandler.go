package main

import (
	"net/http"

	"github.com/kevingentile/kevingentile.com/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	page := web.Page{
		Title: "Portfolio Home",
		Class: "homepage",
	}

	banner := web.PageHeading{
		Header: "Welcome to my portfolio!",
	}

	carousel := web.Carousel{
		Articles: []web.ReelArticle{
			{
				Title:     "kevingentile.com",
				Body:      "This website! Built using Go and the HTML5UP site template.",
				ImageLink: "https://github.com/kevingentile/kevingentile.com",
				PicPath:   "../images/pic01.jpg",
			},
			{
				Title:     "GoLinks",
				Body:      "A blockchain library and filesystem integrity tool built in Go",
				ImageLink: "https://github.com/govice/golinks",
				PicPath:   "../images/pic04.jpg",
			},
			{
				Title:     "Golinks Daemon",
				Body:      "Prototype for maintaining an internal P2P blockchain with Golinks",
				ImageLink: "https://github.com/govice/golinks-daemon",
				PicPath:   "../images/pic04.jpg",
			},
			{
				Title:     "Fortnite Tracker API Driver",
				Body:      "Go driver for the Fortnite Tracker API",
				ImageLink: "https://github.com/kevingentile/fortnite-tracker",
				PicPath:   "../images/pic01.jpg",
			},
			{
				Title:     "Fornite Statistics Widget for Open Broadcast Studio ",
				Body:      "Module for displaying statistics for use in live broadcasts",
				ImageLink: "http://www.kevingentile.com/fortnite.html",
				PicPath:   "../images/pic01.jpg",
			},
			{
				Title:     "govice.org",
				Body:      "Landing page for the Govice organization",
				ImageLink: "https://github.com/govice/govice.org",
				PicPath:   "../images/pic04.jpg",
			},
		},
	}
	data := web.Data{"page": page, "carousel": carousel, "nav": Nav, "footer": Footer, "banner": banner, "scripts": ScriptsCommon}
	renderTemplate(w, "index", data)

}

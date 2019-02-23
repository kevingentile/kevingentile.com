package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var limiter <-chan time.Time

func main() {
	log.Println("keivngentile.com Web Start PORT: " + os.Getenv("PORT"))
	router := mux.NewRouter() // Create gorilla router

	limiter = time.Tick(time.Second * 3)

	// "http://kevingentile.com/index.html"
	router.HandleFunc("/", makeHandler(IndexHandler))
	router.HandleFunc("/index.html", makeHandler(IndexHandler))
	// "http://kevingentile.com/contact.html"
	router.HandleFunc("/contact.html", makeHandler(ContactHandler))
	// "http://kevingentile.com/links.html"
	router.HandleFunc("/links.html", makeHandler(LinksHandler))
	// "http://kevingentile.com/obs/laughingcabbage" JSON // TODO rate limit this handler
	router.HandleFunc("/obs/laughingcabbage", rateLimit(handleFortniteData, &limiter))
	// "http://kevingentile.com/assets/*"
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// "http://kevingentile.com/images/*"
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Println(templates.DefinedTemplates())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by environment variable PORT

}

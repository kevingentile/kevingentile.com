package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("keivngentile.com Web Start")
	router := mux.NewRouter() // Create gorilla router

	// "http://kevingentile.com/index.html"
	router.HandleFunc("/", makeHandler(IndexHandler))
	router.HandleFunc("/index.html", makeHandler(IndexHandler))
	// "http://kevingentile.com/contact.html"
	router.HandleFunc("/contact.html", makeHandler(ContactHandler))
	// "http://kevingentile.com/links.html"
	router.HandleFunc("/links.html", makeHandler(LinksHandler))
	// "http://kevingentile.com/assets/*"
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// "http://kevingentile.com/images/*"
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	//log.Println(templates.DefinedTemplates())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by environment variable PORT

}

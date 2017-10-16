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
	// "http://kevingentile.com/left-sidebar.html"
	router.HandleFunc("/left-sidebar.html", makeHandler(LeftHandler))
	// "http://kevingentile.com/right-sidebar.html"
	router.HandleFunc("/right-sidebar.html", makeHandler(RightHandler))
	// "http://kevingentile.com/no-sidebar.html"
	router.HandleFunc("/no-sidebar.html", makeHandler(NosideHandler))
	// "http://kevingentile.com/assets/*"
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //TODO this seems excessive...
	// "http://kevingentile.com/images/*"
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	//fmt.Println(templates.DefinedTemplates())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by environment variable PORT

}

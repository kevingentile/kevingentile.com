package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Web Start")
	router := mux.NewRouter() // Create gorilla router

	// "http://kevingentile.com/index.html"
	router.HandleFunc("/", makeHandler(IndexHandler))
	// "http://kevingentile.com/assets/*"
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) //TODO this seems excessive...
	// "http://kevingentile.com/images/*"
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by environment variable PORT

}

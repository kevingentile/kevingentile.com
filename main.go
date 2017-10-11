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
	router.HandleFunc("/", makeHandler(IndexHandler))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.PathPrefix("/images").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by env variable PORT

}

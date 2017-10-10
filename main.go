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
	router := mux.NewRouter()                                     // Create gorilla router
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router)) // Listen on port defined by env variable PORT

}

//TODO the string doesn't make sense if you can't pass it a route name...
func makeHandler(handle func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, "") //TODO this is definitely not correct
	}
}

func registerRootHandlers(router *mux.Router) {
	router.HandleFunc("/", makeHandler(indexHandler))
}

package main

import "net/http"

//makeHandler is a wrapper for all handler functions
func makeHandler(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handle(w, r)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func LeftHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "left-sidebar")
}

func RightHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "right-sidebar")
}

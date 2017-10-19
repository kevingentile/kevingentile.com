package main

import (
	"net/http"
)

//Data is used to pass various child structures from a handler for templating
type Data map[string]interface{}

//makeHandler is a wrapper for all handler functions
func makeHandler(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handle(w, r)
	}
}

func LeftHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "left-sidebar", nil)
}

func RightHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "right-sidebar", nil)
}

func NosideHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "no-sidebar", nil)
}

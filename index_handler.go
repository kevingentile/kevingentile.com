package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, title)
}

package main

import (
	"net/http"
	"time"
)

//makeHandler is a wrapper for all handler functions
func makeHandler(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(w, r)
	}
}

func rateLimit(handle func(http.ResponseWriter, *http.Request), limiter *<-chan time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		<-*limiter
		handle(w, r)
	}
}

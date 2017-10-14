package main

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("/pages/" + filename) // load from pages folder
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

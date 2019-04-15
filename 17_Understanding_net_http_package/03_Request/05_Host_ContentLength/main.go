package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

//The type handler serves http reponse writer and pointer to a request !!
//Any type with this method is a handler

type marathon int

func (nip marathon) ServeHTTP(w http.ResponseWriter, requester *http.Request) {
	err := requester.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		ContentLength int64
	}{
		requester.Method,
		requester.URL,
		requester.Form,
		requester.Header,
		requester.ContentLength,
	}
	//Gives back the form, which is a map with key and value
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var knowledge marathon
	http.ListenAndServe(":8080", knowledge)
}

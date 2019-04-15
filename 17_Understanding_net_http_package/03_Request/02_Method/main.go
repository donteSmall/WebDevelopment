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

func (nip marathon) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
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

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//The type handler serves http reponse writer and pointer to a request !!
//Any type with this method is a handler

type marathon int

func (nip marathon) ServeHTTP(w http.ResponseWriter, requester *http.Request) {
	err := requester.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//Gives back the form, which is a map with key and value
	tpl.ExecuteTemplate(w, "index.gohtml", requester.Form)

	fmt.Println("Any code you want to the console !")
	fmt.Fprintln(w, "Any code you want in this func")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var knowledge marathon
	http.ListenAndServe(":8080", knowledge)
}

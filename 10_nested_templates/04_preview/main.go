package main

import (
	"log"
	"net/http"
	"text/template"
)

// Syntax reads package.Type||Function
var temp *template.Template

func init() {
	// Must gaves back the pointer template
	temp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
func index(res http.ResponseWriter, req *http.Request) {
	err := temp.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)

	}
}

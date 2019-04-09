package main

import (
	"log"
	"os"
	"text/template"
)

// Syntax reads package.Type||Function
var temp *template.Template

func init() {
	// Must gaves back the pointer template
	temp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := temp.ExecuteTemplate(os.Stdout, "temp.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

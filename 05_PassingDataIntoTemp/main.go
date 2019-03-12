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
	temp = template.Must(template.ParseFiles("temp.gohtml"))
}

func main() {

	// Here we pass in the data (42) to the template
	err := temp.ExecuteTemplate(os.Stdout, "temp.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

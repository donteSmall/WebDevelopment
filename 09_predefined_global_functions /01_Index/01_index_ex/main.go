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

	inputVals := []string{"zero", "one", "two", "three", "four", "five"}
	err := temp.Execute(os.Stdout, inputVals)
	if err != nil {
		log.Fatalln(err)
	}
}

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

	/*
	   Here we pass in a string to the template.

	   We assign values to variables in our temp.gohtml like :
	   ASSIGN
	   {{$wisdom := .}}
	   USE

	   {{$wisdom}}
	*/
	err := temp.ExecuteTemplate(os.Stdout, "temp.gohtml", `Release the dragon, ohhhhh yea!`)
	if err != nil {
		log.Fatalln(err)
	}
}

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
	// Mirror Unix piping concept with range where the output one thing is the input of another
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "Martin Lurther King",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}
	err := temp.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

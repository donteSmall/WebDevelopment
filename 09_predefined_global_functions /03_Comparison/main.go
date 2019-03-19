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

	scoreboard := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := temp.Execute(os.Stdout, scoreboard)
	if err != nil {
		log.Fatalln(err)
	}
}

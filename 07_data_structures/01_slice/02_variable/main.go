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
	sages := []string{"Malcom X", "Elijah muhammad", "Farrakhan", "Muhammad Ali"}
	err := temp.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

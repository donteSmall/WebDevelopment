package main

import (
	"log"
	"os"
	"text/template"
)

// Syntax reads package.Type||Function
var temp *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// Must gaves back the pointer template
	temp = template.Must(template.ParseFiles("temp.gohtml"))
}

func main() {
	// Mirror Unix piping concept with range where the output one thing is the input of another
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	buddha := sage{
		Name:  "Malcom X",
		Motto: "By any means necessary",
	}

	Muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{mlk, gandhi, buddha, Muhammad}

	err := temp.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

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

type users struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := users{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := users{
		Name:  "Jay-Z",
		Motto: "It's the Rock !",
		Admin: false,
	}
	u3 := users{
		Name:  "Biggie ",
		Motto: "It was all a dream",
		Admin: true,
	}
	u4 := users{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}
	inputVals := []users{u1, u2, u3, u4}

	err := temp.Execute(os.Stdout, inputVals)
	if err != nil {
		log.Fatalln(err)
	}
}

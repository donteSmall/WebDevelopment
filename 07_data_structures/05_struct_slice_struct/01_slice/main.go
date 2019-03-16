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
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom []sage
	Wheels []car
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

	honda := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	bmw := car{
		Manufacturer: "BMW",
		Model:        "525i",
		Doors:        4,
	}
	sages := []sage{mlk, gandhi, buddha, Muhammad}
	cars := []car{honda, bmw}

	datacontainer := items{
		Wisdom: sages,
		Wheels: cars,
	}

	err := temp.Execute(os.Stdout, datacontainer)
	if err != nil {
		log.Fatalln(err)
	}
}

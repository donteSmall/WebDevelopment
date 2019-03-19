package main

import (
	"log"
	"os"
	"strings"
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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// Must gaves back the pointer template

	temp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.gohtml"))
}

func firstThree(trimer string) string {
	trimer = strings.TrimSpace(trimer)
	if len(trimer) >= 3 {
		trimer = trimer[:3]
	}
	return trimer
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
	sages := []sage{mlk, gandhi, buddha}
	cars := []car{honda, bmw}

	datacontainer := struct {
		Wisdom []sage
		Wheels []car
	}{
		Wisdom: sages,
		Wheels: cars,
	}

	err := temp.ExecuteTemplate(os.Stdout, "tmp.gohtml", datacontainer)
	if err != nil {
		log.Fatalln(err)
	}
}

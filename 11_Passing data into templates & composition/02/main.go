package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToLove bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	p1 := doubleZero{
		person{
			Name: "Nipsey Hussel",
			Age:  33,
		},
		true,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}

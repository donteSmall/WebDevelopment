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

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	p1 := person{
		Name: "Nipsey Hussel",
		Age:  33,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}

//Best Practice:
//Call functions in templates for formatting only;not processing or data access
// The main reasons you don't want to do any data  processing in your template:
//(1) Seperation of concerns
//(2) If you're using a functions more than once in a template,
//    the server needs to do the processing more than once.
//    though the standard library might cache processing

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
	buddha := sage{
		Name:  "Micheal Jordan",
		Motto: `Ive failed over and over and over again in my life and that is why I succeed.Talent wins games, but teamwork and intelligence wins championships.I can accept failure, everyone fails at something. But I can't accept not trying.`,
	}
	err := temp.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

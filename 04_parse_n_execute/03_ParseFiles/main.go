package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tplcontainer, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tplcontainer.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tplcontainer, err = tplcontainer.ParseFiles("two.gohtml", "vespa.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tplcontainer.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tplcontainer.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tplcontainer.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tplcontainer.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parse file, takes the filename
	temp, err := template.ParseFiles("temp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

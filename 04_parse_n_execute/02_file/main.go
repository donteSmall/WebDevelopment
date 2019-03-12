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
	// Creates a file and store the contents of temp.gohtml in it !
	newfile, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer newfile.Close()
	err = temp.Execute(newfile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

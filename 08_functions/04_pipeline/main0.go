package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var temp *template.Template

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2016")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := temp.ExecuteTemplate(os.Stdout, "tmp.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

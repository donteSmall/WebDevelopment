package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"flodouble":     double,
	"flosquare":     square,
	"flosquareroot": sqRoot,
}

func main() {

	err := temp.ExecuteTemplate(os.Stdout, "tmp.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

//The type handler serves http reponse writer and pointer to a request !!
//Any type with this method is a handler

type marathon int

func (nip marathon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want to the console !")
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var knowledge marathon
	http.ListenAndServe(":8080", knowledge)
}

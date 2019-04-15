package main

import (
	"fmt"
	"net/http"
)

type Rocnation int

func (m Rocnation) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Jay-z", "this is from Hov")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var hov Rocnation
	http.ListenAndServe(":8080", hov)
}

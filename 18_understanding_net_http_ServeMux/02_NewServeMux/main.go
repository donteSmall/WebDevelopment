package main

import (
	"io"
	"net/http"
)

type rolling60s int

func (Nip rolling60s) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, " Stay postive, Stay pure,work hard, and show love to your people !")
}

type big int

func (wallace big) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, " Spread love its the Crooklyn Way!")
}

func main() {
	var westcoast rolling60s
	var eastcoast big
	mux := http.NewServeMux()
	//path: /dog/ -> this would be able to catch -> dog/something
	mux.Handle("/dog/", westcoast)
	//path: /cat -> this would  not be able to catch because no forward slash-> cat
	mux.Handle("/cat", eastcoast)

	http.ListenAndServe(":8080", mux)
}

package main

import (
	"io"
	"net/http"
)

type badboy int

func (bigge badboy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/NY":
		io.WriteString(res, "Spread love it's the Brooklyn Way, Biggie !")
	case "/Cali":
		io.WriteString(res, "Show love to your people, Neighborhood Nip !")

	}
}
func main() {
	var diddy badboy
	http.ListenAndServe(":8080", diddy)
}

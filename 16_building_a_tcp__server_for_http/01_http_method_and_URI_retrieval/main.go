package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	//When are in the request line when == 0
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			request_line := strings.Fields(ln)[0] // Get  METHOD
			get_urI := strings.Fields(ln)[1]      // Get  URI
			fmt.Println("***METHOD", request_line)
			fmt.Println("***URI", get_urI)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html>
  <html lang="en">
  <head><meta
   charset="UTF-8"><title></title>
   </head>
   <body><strong> IN NIP WE TRUST, ON God ! </strong></body>
   </html>`
	//Status line|version|reason phrase
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	//Header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	//Blank line, char return line feed
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

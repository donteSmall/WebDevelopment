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
	//respond(conn)
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
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	request_line := strings.Fields(ln)[0] // Get  METHOD
	get_urI := strings.Fields(ln)[1]      // Get  URI
	fmt.Println("***METHOD", request_line)
	fmt.Println("***URI", get_urI)

	if request_line == "GET" && get_urI == "/" {
		index(conn)
	}
	if request_line == "GET" && get_urI == "/about" {
		about(conn)
	}

	if request_line == "GET" && get_urI == "/contact" {
		contact(conn)
	}
	if request_line == "GET" && get_urI == "/apply" {
		apply(conn)
	}
	if request_line == "POST" && get_urI == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
		<strong>INDEX</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/contact">contact</a><br>
		<a href="/apply">apply</a><br>
		</body>
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

func about(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en">
	<head><meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	//Status line|version|reason phrase
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	//Header
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	//Blank line, char return line feed
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
		<strong>CONTACT</strong>
		<br>
		<a href="/">index</a>
		<br>
		<a href="/about">about</a>
		<br>
		<a href="/contact">contact</a>
		br>
		<a href="/apply">apply</a>
		<br>
		</body>
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

func apply(conn net.Conn) {

	body := `<!DOCTYPE html>
	<html lang="en">
	<head><meta charset="UTF-8">
	<title> **Title Goes Here**  </title>
	</head>
	<body>
<strong>APPLY</strong>
<br>
<a href="/">index</a>
<br>
<a href="/about">about</a>
<br>
<a href="/contact">contact</a>
<br>
<a href="/apply">apply</a>
<br>

<form method="POST" action="/apply">
<input type="submit" value="appling pressure">
</form>

</body>
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

func applyProcess(conn net.Conn) {

	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>APPLY PROCESS, IN NIP WE TRUST</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply pressure</a><br>
	</body>
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

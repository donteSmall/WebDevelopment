package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		//First thing that happens is we accept, if someone calls
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//Here we start gorountine, passing it to the func handle
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

	}
	defer conn.Close()

	fmt.Println("Does this ever excute")
}

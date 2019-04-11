package main

import (
	"fmt"
	"io"
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
		io.WriteString(conn, "\nLife is a Marthon, not a sprint\n")
		fmt.Fprintln(conn, "How is your day ?")
		fmt.Fprintf(conn, "%v", "Time to double up, three or four times. GOTTA DOUBLE UP!")

		conn.Close()
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	listen , err := net.Listen("tcp",":8084")
	if err!=nil {
		log.Panic(err)
	}

	defer listen.Close()

	for  {
		connection, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(connection, "\nHello from TCP serve\n")
		fmt.Fprintln(connection, "HOW IS YOUR DAY?")
		fmt.Fprintf(connection, "%v","Well,I hope!")

		connection.Close()
	}


	
}

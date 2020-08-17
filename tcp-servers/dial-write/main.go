package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:8084")
	if err != nil {
		log.Panic(err)
	}
	defer connection.Close()

	fmt.Fprintln(connection, "I dialed you!")

}

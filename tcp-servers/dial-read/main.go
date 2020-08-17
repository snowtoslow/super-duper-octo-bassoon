package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	connection, err := net.Dial("tcp", "localhost:8084")
	if err != nil {
		log.Panic(err)
	}
	defer connection.Close()

	bs, err := ioutil.ReadAll(connection)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}

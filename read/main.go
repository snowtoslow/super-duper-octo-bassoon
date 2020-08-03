package main

import (
"bufio"
"fmt"
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
		go handle(connection)
	}



}

func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code got here!")
}


package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp",":8084")

	if err != nil {
		log.Panic(err)
	}

	for  {
		connection, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(connection)
	}
}

func handle(con net.Conn)  {

	err := con.SetDeadline(time.Now().Add(10+time.Second))

	if err != nil {
		log.Panic("Connection time out")
	}

	scanner:= bufio.NewScanner(con)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(con,"You've said: %s\n",line)
	}

	defer con.Close()

	fmt.Println("***CODE GET HERE**")


}

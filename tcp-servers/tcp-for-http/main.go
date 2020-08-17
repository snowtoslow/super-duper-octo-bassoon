package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	listen, err := net.Listen("tcp",":8084")
	if err != nil {
		log.Panic(err)
	}

	defer listen.Close()

	for {
		connection, err := listen.Accept()

		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(connection)
	}

}

func handle(con net.Conn){

	defer con.Close()

	request(con)

	response(con)

}

func request(conn net.Conn){
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			m := strings.Fields(line)[0] // method
			u := strings.Fields(line)[1] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if line==" " {
			break
		}
		i++
	}
}

func response(conn net.Conn){
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}



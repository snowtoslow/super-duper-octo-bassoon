package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	listen, err := net.Listen("tcp", ":8084")

	if err != nil {
		log.Panic(err.Error())
	}

	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(connection)
	}
}

func handle(conn net.Conn){

	defer conn.Close()

	request(conn)

}

func request(conn net.Conn){
	i :=0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		line := scanner.Text()

		fmt.Println("LINE:",line)

		if i == 0 {

			mux(conn, line)
		}

		if line =="" {
			break
		}
		i++
	}

}

func mux(conn net.Conn, line string)  {

	method := strings.Fields(line)[0]
	uri := strings.Fields(line)[1]

	fmt.Println("***METHOD", method)
	fmt.Println("***URI", uri)

	if method == "GET" && uri == "/about" {
		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>ABOUT</strong></body></html>`

		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}else if method=="GET" && uri=="/how" {
		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>HOW</strong></body></html>`

		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}else {
		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}


}

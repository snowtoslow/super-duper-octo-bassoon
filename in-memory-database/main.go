package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Panic(err)
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

func handle(connection net.Conn) {
	defer connection.Close()

	io.WriteString(connection, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(connection, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(connection, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(connection, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}

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
	if err!=nil {
		log.Fatalln(err)
	}

	defer listen.Close()

	for  {
		connection, err := listen.Accept()
		if err!=nil {
			log.Println(err)
			continue
		}
		go handle(connection)
	}

}

func handle(connection net.Conn){
	scanner := bufio.NewScanner(connection)

	for scanner.Scan(){
		line := strings.ToLower(scanner.Text())
		bs := []byte(line)
		r := rot13(bs)

		fmt.Fprintf(connection, "%s - %s\n\n", line, r)
	}
}

func rot13(bs []byte) []byte  {

	var r13 = make([]byte, len(bs))

	for i, v:= range bs{
		if v<=109 {
			r13[i] = v + 13

		}else{
			r13[i] = v - 13
		}
	}
	return r13
}

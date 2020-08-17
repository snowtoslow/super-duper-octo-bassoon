package main


// handler interface is serveHttp write and pointer to request

import (
	"fmt"
	"net/http"
)


type hotdog int


func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Any code here")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8084",d)
}

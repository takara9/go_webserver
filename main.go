package main

import (
	"fmt"
	"net/http"
	"github.com/takara9/go_util"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	port := ":" + go_util.Config.TcpPort
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}

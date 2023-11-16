package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	socket, err := net.Listen("unix", "/~/conn.sock")
	if err != nil {
		panic(err)
	}

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Request received")
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	m := http.NewServeMux()
	m.HandleFunc("/", h1)

	server := http.Server{
		Handler: m,
	}

	fmt.Println("Starting Server")
	if err := server.Serve(socket); err != nil {
		log.Fatal(err)
	}

}

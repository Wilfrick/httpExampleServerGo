package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Request received")
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	http.HandleFunc("/", h1)
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

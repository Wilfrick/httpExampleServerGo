package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	http.HandleFunc("/", h1)
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

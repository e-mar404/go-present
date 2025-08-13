package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))
	mux.Handle("/", fs)
	fmt.Println("serving files from ./ on :8080")

	err := http.ListenAndServe(":8080", mux) 
	if err != nil {
		fmt.Println("error starting server")
	}
}

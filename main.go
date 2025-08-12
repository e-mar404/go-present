package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	fmt.Println("serving files from ./ on :8080")

	err := http.ListenAndServe(":8080", nil) 
	if err != nil {
		fmt.Println("error starting server")
	}
}

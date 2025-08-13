package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Present struct {
}

func (p *Present) NextSlide(reply *string) error {
	*reply = "going to next slide"
	return nil
}

func main() {
	presentor := new(Present)
	rpc.Register(presentor)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))

	mux.Handle("/", fs)
	mux.Handle(rpc.DefaultRPCPath, rpc.DefaultServer)
	mux.Handle(rpc.DefaultDebugPath, rpc.DefaultServer)

	fmt.Println("serving files from ./ on :8080")

	err := http.ListenAndServe(":8080", mux) 
	if err != nil {
		fmt.Println("error starting server")
	}
}

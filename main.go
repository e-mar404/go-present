package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"reflect"

	"github.com/hashicorp/go-msgpack/codec"
)

type Present struct {
	Path string
}

func (p *Present) SetPath(arg string, reply *string) error {	
	// should prob check to see if path has md files
	p.Path = arg
	*reply = "path has been set to " + arg
	fmt.Printf("SetPath: %v\n", *reply)
	return nil
}

func (p *Present) NextSlide(args any, reply *string) error {
	*reply = "going to next slide"
	fmt.Printf("doing cmd: %v\n", *reply)
	return nil
}

func main() {
	pathSet := make(chan bool)
	presentor := new(Present)

	go func() {
		rpc.Register(presentor)

		listener, err := net.Listen("tcp", ":8081")
		if err != nil {
			fmt.Println("PRC listening error: ", err)
		}
		fmt.Println("MessagePack RPC server running on :8081")

		var mh codec.MsgpackHandle
		mh.MapType = reflect.TypeOf(map[string]any(nil))

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("RPC error accepting connection: ", err)
				continue
			}

			if presentor.Path != "" {
				pathSet <- true
			}

			go rpc.ServeCodec(codec.MsgpackSpecRpc.ServerCodec(conn, &mh))
		}

	}()

	receivedPath := <- pathSet

	go func() {
		if receivedPath {
			http.Handle("/", http.FileServer(http.Dir(presentor.Path)))
			fmt.Println("serving files from ", presentor.Path, " on :8080")
			err := http.ListenAndServe(":8080", nil) 
			if err != nil {
				fmt.Println("error starting server")
			}
		}
	}()

	select {}
}

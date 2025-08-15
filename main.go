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
}

func (p *Present) NextSlide(args any, reply *string) error {
	*reply = "going to next slide"
	fmt.Printf("doing cmd: %v\n", *reply)
	return nil
}

func main() {
	go func() {
		http.Handle("/", http.FileServer(http.Dir("./")))
		fmt.Println("serving files from ./ on :8080")
		err := http.ListenAndServe(":8080", nil) 
		if err != nil {
			fmt.Println("error starting server")
		}
	}()

	go func() {
		presentor := new(Present)
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
			go rpc.ServeCodec(codec.MsgpackSpecRpc.ServerCodec(conn, &mh))
		}

	}()

	select {}
}

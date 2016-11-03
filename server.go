package goadmin

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//StartServer start the admin server which will block the goroutine
func StartServer(addr string) {
	if err := rpc.Register(new(Admin)); err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.Serve(listen, nil); err != nil {
		log.Fatal(err)
	}
}

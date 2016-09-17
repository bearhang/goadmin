package goadmin

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//StartServer start the admin server which will block the goroutine
func StartServer(addr string) {
	admin := new(Admin)
	err := rpc.Register(admin)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	err = http.Serve(listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}

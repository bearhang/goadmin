package goadmin

import (
	"log"
	"net/rpc"
)

// StartClient start the client
// addr is IP/Port address of server
// args is command line
func StartClient(addr string, args []string) {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	req := AdminReq{
		Args: args,
	}
	rsp := AdminRsp{}
	err = client.Call("Admin.Query", &req, &rsp)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(rsp.Result)
}

package goadmin

import (
	"log"
	"net/rpc"
)

func StartClient(addr string, args []string) {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	req := AdminReq{
		args: args,
	}
	rsp := AdminRsp{}
	err = client.Call("Admin/Qeury", req, &rsp)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(rsp.result)
}

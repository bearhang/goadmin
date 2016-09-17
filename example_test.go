package goadmin_test

import (
	"fmt"
	"testing"

	"github.com/bearhang/goadmin"
)

type test1 struct {
	Int    int    `usage:"int usage"`
	String string `usage:"string usage"`
}

func test1Handler(input interface{}) string {
	in := input.(*test1)
	return fmt.Sprint(in)
}

func TestExample(t *testing.T) {
	//addr := "127.0.0.1:9204"
	go goadmin.StartServer(":9204")

	err := goadmin.Register("test1", test1Handler, test1{}, "test1 desc")
	if err != nil {
		t.Fatal(err)
	}

	args := []string{
		"test1", "-int=1", "-string=abc",
	}
	goadmin.StartClient("127.0.0.1:9204", args)
}

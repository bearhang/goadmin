package goadmin_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bearhang/goadmin"
)

type test1 struct {
	Int    int    `usage:"int usage"`
	String string `usage:"string usage"`
}

func test1Handler(input interface{}) (string, error) {
	in := input.(*test1)
	return fmt.Sprint(in), nil
}

func TestExample(t *testing.T) {
	//addr := "127.0.0.1:9204"
	go goadmin.StartServer("127.0.0.1:9204")

	time.Sleep(5 * time.Second)
	err := goadmin.Register("test1", "test1 desc", test1{}, test1Handler)
	if err != nil {
		t.Fatal(err)
	}

	args := []string{
		"test1", "-int=1", "-string=abc",
	}
	goadmin.StartClient("127.0.0.1:9204", args)
}

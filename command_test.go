package goadmin

import (
	"fmt"
	"strings"
	"testing"
)

func executeCommand(args []string) (result string) {
	if len(args) == 0 {
		result = showCommands()
		return
	}

	cmd, err := getCommand(args[0])
	if err != nil {
		result = fmt.Sprintln(err)
		result += showCommands()
		return
	}

	result = cmd.excute(args[1:])
	return
}

func TestParseBool(t *testing.T) {
	type testStruct struct {
		Bool bool
	}
	name := "parseBool"
	desc := "test parse bool"
	args := "parseBool -bool"
	expectStruct := &testStruct{
		Bool: true,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestParseInt(t *testing.T) {
	type testStruct struct {
		Int   int
		Int8  int8
		Int16 int16
		Int32 int32
		Int64 int64
	}
	name := "parseInt"
	desc := "test parse int"
	args := "parseInt -int 1 -int8=8 -int16 -16 -int32=-32"
	expectStruct := &testStruct{
		Int:   1,
		Int8:  8,
		Int16: -16,
		Int32: -32,
		Int64: 0,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestParseUint(t *testing.T) {
	type testStruct struct {
		Uint   uint
		Uint8  uint8
		Uint16 uint16
		Uint32 uint32
		Uint64 uint64
	}
	name := "parseUint"
	desc := "test parse uint"
	args := "parseUint -uint 1 -uint8=8 -uint16 16 -uint32=32"
	expectStruct := &testStruct{
		Uint:   1,
		Uint8:  8,
		Uint16: 16,
		Uint32: 32,
		Uint64: 0,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestParseFloat(t *testing.T) {
	type testStruct struct {
		Float32 float32
		Float64 float64
	}
	name := "parseFloat"
	desc := "test parse float"
	args := "parseFloat -float32 3.1415 -float64=3.1415926535"
	expectStruct := &testStruct{
		Float32: 3.1415,
		Float64: 3.1415926535,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestParseString(t *testing.T) {
	type testStruct struct {
		String string
	}
	name := "parseString"
	desc := "test parse String"
	args := "parseString -string=helloworld"
	expectStruct := &testStruct{
		String: "helloworld",
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestRawArgs(t *testing.T) {
	type testStruct struct {
		Int int
	}
	name := "parseRaw"
	desc := "test parse raw argument name"
	args := "parseRaw -Int=1"
	expectStruct := &testStruct{
		Int: 1,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestLowerArgs(t *testing.T) {
	type testStruct struct {
		Int int
	}
	name := "parseLower"
	desc := "test parse lower case argument name"
	args := "parseLower -int=1"
	expectStruct := &testStruct{
		Int: 1,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

func TestAcronymArgs(t *testing.T) {
	type testStruct struct {
		Int int `acronym:"i"`
	}
	name := "parseAcronym"
	desc := "test parse acronym argument name"
	args := "parseAcronym -i=1"
	expectStruct := &testStruct{
		Int: 1,
	}
	err := Register(name, desc, &testStruct{}, func(input interface{}) (string, error) {
		inputStruct := input.(*testStruct)
		if *inputStruct != *expectStruct {
			t.Errorf("case=%s expect=%v result=%v", name, expectStruct, inputStruct)
		}
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer Unregister(name)

	executeCommand(strings.Split(args, " "))
}

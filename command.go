package goadmin

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
)

type argument struct {
	name    string
	value   reflect.Value
	usage   string
	argKind reflect.Kind
}

type command struct {
	name      string
	desc      string
	handler   func(input interface{}) string
	inputType reflect.Type
	args      []argument
}

var commandMap = make(map[string]command)

func getCommand(name string) (*command, error) {
	cmd, ok := commandMap[name]
	if !ok {
		return nil, errors.New("command not found")
	}
	return &cmd, nil
}

func showCommands() string {
	buf := "SYNOPSIS:\n"
	for _, cmd := range commandMap {
		buf += fmt.Sprintf("\t%s\n", cmd.synopsis())
	}
	return buf
}

func setFlag(fs *flag.FlagSet, arg *argument) interface{} {
	switch arg.argKind {
	case reflect.Bool:
		return fs.Bool(arg.name, arg.value.Bool(), arg.usage)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fs.Int64(arg.name, arg.value.Int(), arg.usage)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fs.Uint64(arg.name, arg.value.Uint(), arg.usage)
	case reflect.Float32, reflect.Float64:
		return fs.Float64(arg.name, arg.value.Float(), arg.usage)
	case reflect.String:
		return fs.String(arg.name, arg.value.String(), arg.usage)
	default:
		panic(fmt.Sprintf("unsupport kind %s", arg.argKind.String()))
	}
}

func (cmd *command) excute(args []string) string {
	fs := flag.NewFlagSet(cmd.name, flag.ContinueOnError)
	argsPrt := make(map[string]interface{})
	for _, arg := range cmd.args {
		argsPrt[arg.name] = setFlag(fs, &arg)
	}

	err := fs.Parse(args)
	if err != nil {
		return ""
	}

	inputValue := reflect.New(cmd.inputType)
	for name, ptr := range argsPrt {
		field := inputValue.FieldByName(name)
		field.Set(reflect.ValueOf(ptr))
	}

	return cmd.handler(inputValue.Interface())
}

func (cmd *command) synopsis() string {
	var buf string
	buf += fmt.Sprintf("%s\t%s\n", cmd.name, cmd.desc)
	return buf
}

func (cmd *command) description() string {
	return ""
}

// Register register command
// command should be registered at server end
func Register(name string, handler func(input interface{}) string,
	input interface{}, desc string) error {
	if _, ok := commandMap[name]; ok {
		return fmt.Errorf("command %s had been registered", name)
	}

	var cmd command
	cmd.args = make([]argument, 0)
	cmd.name = name
	cmd.desc = desc
	cmd.handler = handler
	cmd.inputType = reflect.TypeOf(input)

	commandMap[name] = cmd
	return nil
}

// Unregister unregiseter command
func Unregister(name string) {
	delete(commandMap, name)
}

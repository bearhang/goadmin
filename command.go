package goadmin

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strings"
)

type argument struct {
	name     string
	acronym  string
	usage    string
	argKind  reflect.Kind
	argValue reflect.Value
}

type command struct {
	name      string
	desc      string
	handler   func(input interface{}) (string, error)
	inputType reflect.Type
	args      map[string]argument
}

var commandMap = make(map[string]command)

func (cmd *command) excute(args []string) string {
	fs := flag.NewFlagSet(cmd.name, flag.ContinueOnError)
	argsPtr := make(map[string]interface{})
	inputValue := reflect.New(cmd.inputType)

	for name, arg := range cmd.args {
		argsPtr[name] = setFlag(fs, &arg)
	}

	if err := fs.Parse(args); err != nil {
		return err.Error()
	}

	for name, ptr := range argsPtr {
		field := inputValue.Elem().FieldByName(name)
		field.Set(reflect.ValueOf(ptr).Elem().Convert(field.Type()))
	}

	ret, err := cmd.handler(inputValue.Interface())
	if err != nil {
		ret = fmt.Sprintln(err)
		ret += "SYNOPSIS\n"
		ret += cmd.synopsis()
		ret += "\nDESCRIPTION\n"
		ret += cmd.description()
	}
	return ret
}

func (cmd *command) synopsis() string {
	var buf string
	buf += fmt.Sprintf("\t%s\t%s\n", cmd.name, cmd.desc)
	return buf
}

func (cmd *command) description() string {
	var buf string
	for _, arg := range cmd.args {
		buf += fmt.Sprintf("\t-%s", arg.name)
		if arg.acronym != "" {
			buf += fmt.Sprintf("|-%s", arg.acronym)
		}
		buf += fmt.Sprintf("\t%s\n", arg.usage)
	}
	return buf
}

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
		var value bool
		fs.BoolVar(&value, arg.name, arg.argValue.Bool(), arg.usage)
		fs.BoolVar(&value, strings.ToLower(arg.name), arg.argValue.Bool(), arg.usage)
		if arg.acronym != "" {
			fs.BoolVar(&value, arg.acronym, arg.argValue.Bool(), arg.usage)
		}
		return &value
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var value int64
		fs.Int64Var(&value, arg.name, arg.argValue.Int(), arg.usage)
		fs.Int64Var(&value, strings.ToLower(arg.name), arg.argValue.Int(), arg.usage)
		if arg.acronym != "" {
			fs.Int64Var(&value, arg.acronym, arg.argValue.Int(), arg.usage)
		}
		return &value
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var value uint64
		fs.Uint64Var(&value, arg.name, arg.argValue.Uint(), arg.usage)
		fs.Uint64Var(&value, strings.ToLower(arg.name), arg.argValue.Uint(), arg.usage)
		if arg.acronym != "" {
			fs.Uint64Var(&value, arg.acronym, arg.argValue.Uint(), arg.usage)
		}
		return &value
	case reflect.Float32, reflect.Float64:
		var value float64
		fs.Float64Var(&value, arg.name, arg.argValue.Float(), arg.usage)
		fs.Float64Var(&value, strings.ToLower(arg.name), arg.argValue.Float(), arg.usage)
		if arg.acronym != "" {
			fs.Float64Var(&value, arg.acronym, arg.argValue.Float(), arg.usage)
		}
		return &value
	case reflect.String:
		var value string
		fs.StringVar(&value, arg.name, arg.argValue.String(), arg.usage)
		fs.StringVar(&value, strings.ToLower(arg.name), arg.argValue.String(), arg.usage)
		if arg.acronym != "" {
			fs.StringVar(&value, arg.acronym, arg.argValue.String(), arg.usage)
		}
		return &value
	default:
		panic(fmt.Sprintf("unsupport kind %s", arg.argKind))
	}
}

func validArgument(arg *argument) error {
	switch arg.argKind {
	case reflect.Bool:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Float32, reflect.Float64:
	case reflect.String:
	default:
		return fmt.Errorf("unsupport argument type %s", arg.argKind.String())
	}
	return nil
}

// Register register command
// command should be registered at server end
func Register(name string, handler func(input interface{}) (string, error),
	input interface{}, desc string) error {
	if _, ok := commandMap[name]; ok {
		return fmt.Errorf("command %s had been registered", name)
	}

	var cmd command
	cmd.args = make(map[string]argument)
	cmd.name = name
	cmd.desc = desc
	cmd.handler = handler
	cmd.inputType = reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	if cmd.inputType.Kind() == reflect.Ptr {
		cmd.inputType = cmd.inputType.Elem()
		inputValue = inputValue.Elem()
	}
	if cmd.inputType.Kind() != reflect.Struct {
		return errors.New("type of struct should be struct or pointer to struct")
	}

	for i := 0; i < cmd.inputType.NumField(); i++ {
		var arg argument
		field := cmd.inputType.Field(i)
		arg.name = field.Name
		arg.acronym = field.Tag.Get("acronym")
		arg.usage = field.Tag.Get("usage")
		arg.argKind = field.Type.Kind()
		arg.argValue = inputValue.Field(i)
		if err := validArgument(&arg); err != nil {
			return err
		}
		cmd.args[arg.name] = arg
	}
	commandMap[name] = cmd
	return nil
}

// Unregister unregiseter command
func Unregister(name string) {
	delete(commandMap, name)
}

package goadmin

import "fmt"

type AdminReq struct {
	args []string
}

type AdminRsp struct {
	result string
}

type Admin int

func (a *Admin) Query(req *AdminReq, rsp *AdminRsp) error {
	if len(req.args) == 0 {
		rsp.result = showCommands()
		return nil
	}

	cmd, err := getCommand(req.args[0])
	if err != nil {
		rsp.result = fmt.Sprintln(err)
		rsp.result += showCommands()
		return nil
	}

	rsp.result = cmd.excute(req.args[1:])
	return nil
}

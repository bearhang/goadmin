package goadmin

import "fmt"

// AdminReq admin require struct
type AdminReq struct {
	Args []string
}

// AdminRsp admin response struct
type AdminRsp struct {
	Result string
}

// Admin admin service
type Admin int

// Query admin service's handler'
func (a *Admin) Query(req *AdminReq, rsp *AdminRsp) error {
	if len(req.Args) == 0 {
		rsp.Result = showCommands()
		return nil
	}

	cmd, err := getCommand(req.Args[0])
	if err != nil {
		rsp.Result = fmt.Sprintln(err)
		rsp.Result += "\n" + showCommands()
		return nil
	}

	rsp.Result = cmd.excute(req.Args[1:])
	return nil
}

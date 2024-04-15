package process

import (
	"awesomeProject/common"
	"awesomeProject/server/utils"
	"net"
)

type User struct {
	c net.Conn
}

func (u *User) Login(req common.LoginReq) {
	s := utils.Socket{
		C: u.c,
	}
	if req.Username != "haven" || req.Password != "123" {
		s.WriteData(&common.LoginRes{
			Token: "12oj4019nnq",
		}, 200, "success")
	} else {
		s.WriteData(&common.LoginRes{
			Token: "",
		}, 401, "error")
	}
}

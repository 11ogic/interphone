package main

import (
	"awesomeProject/common"
	"awesomeProject/server/process"
	"awesomeProject/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type processor struct {
	Conn net.Conn
}

func (p *processor) receive() (err error) {
	for {
		s := &utils.Socket{
			C: p.Conn,
		}
		req, err := s.ReadData()
		if err != nil {
			if err == io.EOF {
				return errors.New("disconnect from client")
			}
			fmt.Println("err:", err)
			return err
		}
		err = p.dispatch(req)
	}
	return
}

func (p *processor) dispatch(req *common.RequestType) (err error) {
	switch req.Type {
	case common.SMS:
		fmt.Println("enter SMS")
	case common.User:
		fmt.Println("enter User")
		u := process.User{
			C: p.Conn,
		}
		data := &common.LoginReq{}
		err = json.Unmarshal([]byte(req.Data), data)
		if err != nil {
			return
		}
		u.Login(data)
	default:
		fmt.Println("unknown type")
	}
	return
}

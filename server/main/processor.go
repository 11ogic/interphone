package main

import (
	"awesomeProject/common"
	"awesomeProject/server/process"
	"awesomeProject/server/utils"
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
		fmt.Printf("%+v", req.Data)
		u := process.User{
			c: p.Conn,
		}
	default:
		fmt.Println("unknown type")
	}
	return
}

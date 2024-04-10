package main

import (
	"awesomeProject/server/common"
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
		m := &utils.Message{
			C: p.Conn,
		}
		req, err := m.ReadData()
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
	default:
		fmt.Println("unknown type")
	}
	return
}

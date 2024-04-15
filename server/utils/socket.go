package utils

import (
	"awesomeProject/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Socket struct {
	C   net.Conn
	buf [8192]byte
}

func (s *Socket) ReadData() (req *common.RequestType, err error) {
	n, err := s.C.Read(s.buf[:4])
	if n != 4 || err != nil {
		err = errors.New("read fail")
	}
	size := binary.BigEndian.Uint32(s.buf[:4])
	n, err = s.C.Read(s.buf[:size])
	if err != nil || size != uint32(n) {
		err = errors.New("read fail")
	}
	err = json.Unmarshal(s.buf[:size], req)
	if err != nil {
		fmt.Println(string(s.buf[:n]))
		return
	}
	fmt.Printf("size = %d; content = %+v \n", size, req)
	return
}

func (s *Socket) WriteData(inter interface{}, code uint32, msg string) (err error) {
	data, err := json.Marshal(inter)
	if err != nil {
		return errors.New("marshal fail")
	}
	request := &common.ResponseType{Data: string(data), Code: code, Msg: msg}
	sendData, err := json.Marshal(request)
	if err != nil {
		return errors.New("marshal fail")
	}
	size := uint32(len(sendData))
	sizeBuf := [4]byte{}
	binary.BigEndian.PutUint32(sizeBuf[:4], size)
	n, err := s.C.Write(sizeBuf[:])
	if n != 4 || err != nil {
		return errors.New("failed to write size")
	}
	n, err = s.C.Write(sendData)
	if n != int(size) || err != nil {
		return errors.New("failed to write sendData")
	}

	return
}

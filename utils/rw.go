package utils

import (
	"awesomeProject/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Message struct {
	C   net.Conn
	buf [8192]byte
}

func (m *Message) ReadData() (req *common.RequestType, err error) {
	n, err := m.C.Read(m.buf[:4])
	if n != 4 || err != nil {
		err = errors.New("read fail")
	}
	size := binary.BigEndian.Uint32(m.buf[:4])
	n, err = m.C.Read(m.buf[:size])
	if err != nil || size != uint32(n) {
		err = errors.New("read fail")
	}
	req = &common.RequestType{}
	err = json.Unmarshal(m.buf[:size], req)
	if err != nil {
		fmt.Println(string(m.buf[:n]))
		return
	}
	fmt.Printf("size = %d; content = %+v \n", size, req)
	return
}

func (m *Message) WriteData(inter interface{}) (err error) {
	data, err := json.Marshal(inter)
	if err != nil {
		return errors.New("marshal fail")
	}
	request := &common.ResponseType{Data: string(data), Code: 200, Msg: "success"}
	sendData, err := json.Marshal(request)
	if err != nil {
		return errors.New("marshal fail")
	}
	size := uint32(len(sendData))
	sizeBuf := [4]byte{}
	binary.BigEndian.PutUint32(sizeBuf[:4], size)
	n, err := m.C.Write(sizeBuf[:])
	if n != 4 || err != nil {
		fmt.Println(n)
		fmt.Println(err)
		return errors.New("failed to write size")
	}
	n, err = m.C.Write(sendData)
	if n != int(size) || err != nil {
		return errors.New("failed to write sendData")
	}

	return
}

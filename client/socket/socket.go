package socket

import (
	"awesomeProject/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Fetch struct {
	C net.Conn
}

func (f *Fetch) Write(typ string, data interface{}) (err error) {
	d, err := json.Marshal(data)
	if err != nil {
		return
	}
	req := common.RequestType{
		Data: string(d),
		Type: typ,
	}
	r, err := json.Marshal(req)
	if err != nil {
		return
	}
	size := uint32(len(r))
	sizeBytes := [4]byte{}
	binary.BigEndian.PutUint32(sizeBytes[:4], size)
	n, err := f.C.Write(sizeBytes[:])
	if n != 4 || err != nil {
		fmt.Println("fetch failed")
		return
	}
	n, err = f.C.Write(r)
	return
}

func (f *Fetch) Read() {

}

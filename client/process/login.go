package process

import (
	"awesomeProject/client/socket"
	"awesomeProject/common"
	"errors"
	"fmt"
	"net"
	"strings"
)

func Login(n net.Conn) (err error) {
	username := ""
	password := ""

	fmt.Println("enter your username")
	fmt.Scanf("%v", &username)
	fmt.Println("enter your password")
	fmt.Scanf("%v", &password)

	if strings.Trim(username, "") == "" || strings.Trim(password, "") == "" {
		err = errors.New("validate failed")
		return
	} else {
		f := socket.Fetch{C: n}
		f.Write(common.User, common.LoginReq{Username: username, Password: password})
		res, err := f.Read()
		if err != nil {
			return errors.New("read failed")
		}
		fmt.Printf("%+v", res)
	}
	return
}

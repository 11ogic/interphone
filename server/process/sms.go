package process

import "net"

type ChatRoom struct {
	c net.Conn
}

func (c *ChatRoom) Join() {

}

func (c *ChatRoom) receive() {}

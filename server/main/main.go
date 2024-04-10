package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen fail")
		return
	}
	fmt.Println("link start")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept fail")
		}
		go func() {
			defer c.Close()
			p := &processor{
				c,
			}
			err := p.receive()
			if err != nil {
				fmt.Println("closed")
				return
			}
		}()
	}
}

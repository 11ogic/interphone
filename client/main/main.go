package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	state uint8
)

func main() {
	_, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("link failure")
		return
	}
	prologue := [...]string{
		"--- welcome to my chat room, please select the options below ~ ---",
		"1. login",
		"2. registry",
		"3. quit",
		"------------------------------------------------------------------",
	}

	for {
		for _, val := range prologue {
			fmt.Println(val)
		}
		fmt.Scanf("%d", &state)
		switch state {
		case 1:
			var (
				username string
				password string
			)
			fmt.Println("enter your username")
			fmt.Scanf("%v", &username)
			fmt.Println("enter your password")
			fmt.Scanf("%v", &password)

			if strings.Trim(username, "") == "" || strings.Trim(password, "") == "" {
				continue
			} else {
				fmt.Println("go login")
				break
			}
		case 2:
			var (
				username string
				password string
			)
			fmt.Println("enter your username")
			fmt.Scanf("%v", &username)
			fmt.Println("enter your password")
			fmt.Scanf("%v", &password)

			if strings.Trim(username, "") == "" || strings.Trim(password, "") == "" {
				continue
			} else {
				fmt.Println("go registry")
				break
			}
		case 3:
			os.Exit(1)
		default:
			continue
		}
	}

}

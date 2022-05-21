package checktcp

import (
	"fmt"
	"net"
)

func Checktcp(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(address + "blocked")
		return
	}
	conn.Close()
	fmt.Println(address + " can access")
}

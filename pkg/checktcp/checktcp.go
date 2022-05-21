package checktcp

import (
	"fmt"
	"net"
	"time"
)

func Checktcp(address string) {
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		fmt.Println(address + "blocked")
		return
	}
	conn.Close()
	fmt.Println(address + " can access")
	return
}

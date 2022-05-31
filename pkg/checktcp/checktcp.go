package checktcp

import (
	"fmt"
	"net"
	"time"
)

func Checktcp(address string, port string) {
	conn, err := net.DialTimeout("tcp", address+":"+port, 2*time.Second)
	if err != nil {
		fmt.Println("[-] tcp", port, "blocked")
		return
	}
	conn.Close()
	fmt.Println("[*] tcp", port, "can access internet")
	return
}

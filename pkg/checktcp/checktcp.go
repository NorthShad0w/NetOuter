package checktcp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ChecktcpN(address string, port string, checknumPtr *int) {
	conn, err := net.DialTimeout("tcp", address+":"+port, 2*time.Second)
	if err != nil {
		//fmt.Println("[-] tcp", port, "blocked")
		return
	}
	conn.Close()
	fmt.Println("[*] tcp", port, "can access internet")
	*checknumPtr++
	return
}

func Checktcp(address string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.DialTimeout("tcp", address+":"+port, 1*time.Second)
	if err != nil {
		fmt.Println("[-] tcp", port, "blocked")
		return
	}
	conn.Close()
	fmt.Println("[*] tcp", port, "can access internet")
	return
}

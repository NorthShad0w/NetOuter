package checktftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Checktftp(target string, wg *sync.WaitGroup) {

	defer wg.Done()

	p := make([]byte, 2048)
	to_sent := []byte{0x00, 0x01, 0x31, 0x2e, 0x74, 0x78, 0x74, 0x00, 0x6e, 0x65, 0x74, 0x61, 0x73, 0x63, 0x69, 0x69}
	conn, err := net.Dial("udp", target+":69")
	conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, string(to_sent))

	_, err = bufio.NewReader(conn).Read(p)
	if p[3] != 0 {
		log.Println("[*] UDP 69  can access the internet")
		return
	}
	log.Println("[-] UDP 69  May blocked")
	conn.Close()
	return
}

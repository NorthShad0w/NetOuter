package checkdns

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

func CheckDirectDNS(wg *sync.WaitGroup) bool {
	defer func() {
		wg.Done()
	}()
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 1,
			}
			return d.DialContext(ctx, "udp", "114.114.114.114:53")
		},
	}
	_, err := r.LookupHost(context.Background(), "www.baidu.com")
	if err != nil {
		log.Println("[-] UDP 53  is blocked")
		return false
	} else {
		log.Println("[*] UDP 53  can access the internet")
		return true
	}

}

func CheckLocalDNS() bool {
	const timeout = 1 * time.Second
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	var r net.Resolver
	resp, err := r.LookupHost(ctx, "www.baidu.com")

	if err != nil {
		log.Println("[-] DNS resolve is blocked")
		return false
	}
	if len(resp) > 0 {
		log.Println("[*] DNS tunnel is allowed")
		return true
	} else {
		log.Println("[-] DNS resolve is blocked")
		return false
	}
}

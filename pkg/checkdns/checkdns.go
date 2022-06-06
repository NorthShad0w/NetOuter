package checkdns

import (
	"context"
	"fmt"
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
		fmt.Println("[-] UDP 53  is blocked")
		return false
	} else {
		fmt.Println("[*] UDP 53  can access the internet")
		return true
	}

}

func CheckLocalDNS() bool {
	resp, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println("[-] DNS resolve is blocked")
		return false
	}
	if len(resp) > 0 {
		fmt.Println("[*] DNS tunnel is allowed")
		return true
	} else {
		fmt.Println("[-] DNS resolve is blocked")
		return false
	}
}

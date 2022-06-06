package checkdns

import (
	"context"
	"fmt"
	"net"
	"time"
)

func CheckDirectDNS() {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", "114.114.114.114:53")
		},
	}
	_, err := r.LookupHost(context.Background(), "www.baidu.com")
	if err != nil {
		fmt.Println("[-] UDP 53  is blocked")
	} else {
		fmt.Println("[*] UDP 53  can access the internet")
	}

}

func CheckLocalDNS() {
	resp, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println("[-] DNS resolve is blocked")
	}
	if len(resp) > 0 {
		fmt.Println("[*] DNS tunnel is allowed")
	} else {
		fmt.Println("[-] DNS resolve is blocked")
	}
}

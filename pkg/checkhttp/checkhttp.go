package checkhttp

import (
	"fmt"
	"net/http"
	"time"
)

var white_domainname [4]string = [...]string{
	"www.baidu.com",
	"www.qq.com",
	"www.gov.cn",
	"mp.weixin.qq.com",
}

func Checkhttp() {

	for _, domain := range white_domainname {
		transport := &http.Transport{}

		timeout := 1 * time.Second

		client := http.Client{
			Timeout:   timeout,
			Transport: transport,
		}

		resp, err := client.Get("http://" + domain + "/")
		if err != nil {
			fmt.Println("[-] http  " + domain + "   was blocked")
		}
		if resp != nil {
			fmt.Println("[*]", domain+"  http protocol is allowed")
		}
	}
	return
}

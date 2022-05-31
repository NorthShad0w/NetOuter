package checkhttp

import (
	"fmt"
	"net/http"
)

var white_domainname [4]string = [...]string{
	"www.baidu.com",
	"www.qq.com",
	"www.gov.cn",
	"mp.weixin.qq.com",
}

func Checkhttp() {

	for _, domain := range white_domainname {
		resp, err := http.Get("http://" + domain + "/")
		if err != nil {
			fmt.Println("[-] http  " + domain + "   was blocked")
		}
		if resp != nil {
			fmt.Println("[*]", domain+"  http protocol is allowed")
		}
	}
}

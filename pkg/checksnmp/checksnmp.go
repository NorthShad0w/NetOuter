package checksnmp

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func Checksnmp(ip string) {
	g.Default.Target = ip
	err := g.Default.Connect()
	if err != nil {
		fmt.Println("[-] UDP 161 is blocked")
	}
	defer g.Default.Conn.Close()
	oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	_, err = g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err != nil {
		fmt.Println("[-] UDP 161 is blocked")
	}
	fmt.Println("[*] UDP 161 can access the internet")
}

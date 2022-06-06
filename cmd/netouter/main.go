package main

import (
	"NetOuter/pkg/checkdns"
	// "NetOuter/pkg/checkhttp"
	"NetOuter/pkg/checkicmp"
	"NetOuter/pkg/checkntp"
	"NetOuter/pkg/checksnmp"
	"NetOuter/pkg/checktcp"
	"NetOuter/pkg/checktftp"
	"flag"
	"fmt"
	"os"
	"sync"
)

var version = "0.1.0"

// flag pointers
var (
	tcpCheckPtr  *string
	snmpCheckPtr *bool
	tftpCheckPtr *bool
	customip     *string
	tcpport      *int
)

var wg sync.WaitGroup

func main() {

	tcpCheckPtr = flag.String("tcp", "", "Check egress for all tcp ports. Example: ./netout -tcp all")
	snmpCheckPtr = flag.Bool("snmp", false, "snmp custom ip check")
	tftpCheckPtr = flag.Bool("tftp", false, "tftp custom ip check")
	customip = flag.String("ip", "1.1.1.1", "custom ip for snmp or tftp")
	tcpport = flag.Int("port", 9999999, "custom tcp port")

	flag.Parse()

	if *tcpport != 9999999 {
		os.Exit(0)
	}

	if *snmpCheckPtr {
		os.Exit(0)
	}
	if *tftpCheckPtr {
		os.Exit(0)
	}

	if *tcpCheckPtr == "all" {
		fmt.Println("[!] All check may take a few minutes to be done, consider using default checking first.")
		fmt.Println("[!] No output means all tcp ports was blocked")
		checktcp.CheckALLtcp()

	} else if *tcpCheckPtr == "test" {

		checktcp.CheckDTCP()

	} else {

		// resp := checkdns.CheckLocalDNS()
		// if resp {
		// 	checkhttp.Checkhttp()
		// }
		wg.Add(7)

		go checkntp.Checkntp(&wg)
		go checksnmp.Checksnmp("116.162.120.19", &wg)
		go checktftp.Checktftp("183.62.177.78", &wg)
		go checkdns.CheckDirectDNS(&wg)
		go checkicmp.Checkicmp(&wg)
		go checktcp.Checktcp("220.181.38.148", "80", &wg)
		go checktcp.Checktcp("220.181.38.148", "443", &wg)
		fmt.Println("[!] Starting default TCP egress check, may take a few minutes to be done.Please Wait patiently.")
		wg.Wait()
		print(1)
		checktcp.CheckDTCP()
		fmt.Println("[!] finished! No tcp output means all tcp ports was blocked")
	}

}

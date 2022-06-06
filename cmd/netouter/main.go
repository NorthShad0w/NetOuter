package main

import (
	"NetOuter/pkg/checkdns"
	"NetOuter/pkg/checkhttp"
	"NetOuter/pkg/checkicmp"
	"NetOuter/pkg/checkntp"
	"NetOuter/pkg/checksnmp"
	"NetOuter/pkg/checktcp"
	"NetOuter/pkg/checktftp"
	"flag"
	"os"
	"strconv"
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

func main() {

	tcpCheckPtr = flag.String("tcp", "", "Check egress for all tcp ports. Example: ./netout -tcp all")
	snmpCheckPtr = flag.Bool("snmp", false, "snmp custom ip check")
	tftpCheckPtr = flag.Bool("tftp", false, "tftp custom ip check")
	customip = flag.String("ip", "1.1.1.1", "custom ip for snmp or tftp")
	tcpport = flag.Int("port", 9999999, "custom tcp port")

	flag.Parse()

	if *tcpport != 9999999 {
		checktcp.Checktcp("45.79.204.144", strconv.Itoa(*tcpport))
		os.Exit(0)
	}

	if *snmpCheckPtr {
		checksnmp.Checksnmp(*customip)
		os.Exit(0)
	}
	if *tftpCheckPtr {
		checktftp.Checktftp(*customip)
		os.Exit(0)
	}

	if *tcpCheckPtr == "all" {
		checktcp.CheckALLtcp()
	} else {
		checkntp.Checkntp()
		checksnmp.Checksnmp("116.162.120.19")
		checktftp.Checktftp("183.62.177.78")
		checkdns.CheckDirectDNS()
		checkdns.CheckLocalDNS()
		checkicmp.Checkicmp()
		checkhttp.Checkhttp()
		checktcp.Checktcp("220.181.38.148", "80")
		checktcp.Checktcp("220.181.38.148", "443")
		checktcp.CheckDTCP()
	}

}

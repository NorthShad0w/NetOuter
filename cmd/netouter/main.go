package main

import (
	"NetOuter/pkg/checkdns"
	"NetOuter/pkg/checkntp"
	"NetOuter/pkg/checkquic"
	"NetOuter/pkg/checksnmp"
	"NetOuter/pkg/checktcp"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error")
		return
	}
	mode := os.Args[1]
	if mode == "d" {
		checkntp.Checkntp()
		checkdns.CheckDirectDNS()
		checkdns.CheckLocalDNS()
		checkquic.Checkquic()
		checktcp.Checktcp("39.156.66.14:80")
		checktcp.Checktcp("39.156.66.14:443")
		checktcp.Checktcp("114.114.114.114:53")
	} else if mode == "tcp" {
		targets_file_path := os.Args[2]
        checktcp.ChecktcpM(targets_file_path)
	} else if mode == "snmp" {
		checksnmp.Checksnmp(os.Args[2])
	}
}

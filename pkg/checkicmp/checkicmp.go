package checkicmp

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func Checkicmp() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("ping.exe", "-n", "1", "114.114.114.114")
		out, err := cmd.CombinedOutput()
		if err != nil {
			print("ping err")
		}
		if strings.Contains(string(out), "TTL") {
			fmt.Println("[*] ICMP allowed")
		} else {
			fmt.Println("[-] ICMP blocked")
		}
	default:
		cmd := exec.Command("ping", "-c", "1", "114.114.114.114")
		out, err := cmd.CombinedOutput()
		if err != nil {
			print("ping err")
		}
		if strings.Contains(string(out), "ttl") {
			fmt.Println("[*] ICMP allowed")
		} else {
			fmt.Println("[-] ICMP blocked")
		}
	}
}

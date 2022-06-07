package checkntp

import (
	"log"
	"sync"

	"github.com/beevik/ntp"
)

func Checkntp(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := ntp.Time("52.231.114.183")
	if err != nil {
		log.Println("[-] NTP protocol(UDP 123) blocked")
		return
	}
	log.Println("[*] UDP 123 can access the internet")
}

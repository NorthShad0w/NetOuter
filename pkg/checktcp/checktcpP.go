package checktcp

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func CheckALLtcp() {
	var slice = make([]int, 65535)
	for i := 1; i <= 65535; i++ {
		slice[i-1] = i
	}
	CheckTCP_port_range(slice)

}

func testHTTPEgress(port int) bool {

	var scheme string
	scheme = "http://"

	url := scheme + "45.79.204.144" + ":" + strconv.Itoa(port)

	transport := &http.Transport{}

	timeout := time.Duration(1000) * time.Second

	client := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	if resp != nil {
		log.Printf("[*] tcp %s can access the internet\n", strconv.Itoa(port))
		return true
	}
	return false
}

type maxedWaitGroup struct {
	current chan int
	wg      sync.WaitGroup
}

func (m *maxedWaitGroup) Add() {
	m.current <- 1
	m.wg.Add(1)
}

func (m *maxedWaitGroup) Done() {
	<-m.current
	m.wg.Done()
}

func (m *maxedWaitGroup) Wait() {
	m.wg.Wait()
}

func CheckDTCP() {
	var ports []int = []int{
		22, 8080, 21, 23, 53, 111, 3389, 4000, 7000, 8000, 8001, 8002, 8003, 8004, 8005, 8006, 8007, 8008, 8009, 8010, 8011, 8012, 8013, 8014, 8015, 8016, 8017, 8018, 8019, 8020, 8021, 8022, 8023, 8024, 8025, 8026, 8027, 8028, 8029, 8030, 8031, 8032, 8033, 8034, 8035, 8036, 8037, 8038, 8039, 8040, 8041, 8042, 8043, 8044, 8045, 8046, 8047, 8048, 8049, 8050, 8051, 8052, 8053, 8054, 8055, 8056, 8057, 8058, 8059, 8060, 8061, 8062, 8063, 8064, 8065, 8066, 8067, 8068, 8069, 8070, 8071, 8072, 8073, 8074, 8075, 8076, 8077, 8078, 8079, 8080, 8081, 8082, 8083, 8084, 8085, 8086, 8087, 8088, 8089, 8090, 8091, 8092, 8093, 8094, 8095, 8096, 8097, 8098, 8099, 8100, 7, 11, 13, 15, 17, 19, 25, 26, 37, 38, 43, 49, 51, 67, 70, 79, 81, 82, 83, 84, 85, 86, 87, 88, 89, 102, 104, 110, 111, 113, 119, 121, 135, 138, 139, 143, 175, 179, 199, 211, 264, 311, 389, 444, 445, 465, 500, 502, 503, 505, 512, 515, 548, 554, 564, 587, 631, 636, 646, 666, 771, 777, 789, 800, 801, 873, 880, 902, 992, 993, 995, 1000, 1022, 1023, 1024, 1025, 1026, 1027, 1080, 1099, 1177, 1194, 1200, 1201, 1234, 1241, 1248, 1260, 1290, 1311, 1344, 1400, 1433, 1471, 1494, 1505, 1515, 1521, 1588, 1720, 1723, 1741, 1777, 1863, 1883, 1911, 1935, 1962, 1967, 1991, 2000, 2001, 2002, 2020, 2022, 2030, 2049, 2080, 2082, 2083, 2086, 2087, 2096, 2121, 2181, 2222, 2223, 2252, 2323, 2332, 2375, 2376, 2379, 2401, 2404, 2424, 2455, 2480, 2501, 2601, 2628, 3000, 3128, 3260, 3288, 3299, 3306, 3307, 3310, 3333, 3388, 3389, 3390, 3460, 3541, 3542, 3689, 3690, 3749, 3780, 4000, 4022, 4040, 4063, 4064, 4369, 4443, 4444, 4505, 4506, 4567, 4664, 4712, 4730, 4782, 4786, 4840, 4848, 4880, 4911, 4949, 5000, 5001, 5002, 5006, 5007, 5009, 5050, 5084, 5222, 5269, 5357, 5400, 5432, 5555, 5560, 5577, 5601, 5631, 5672, 5678, 5800, 5801, 5900, 5901, 5902, 5903, 5938, 5984, 5985, 5986, 6000, 6001, 6068, 6379, 6488, 6560, 6565, 6581, 6588, 6590, 6664, 6665, 6666, 6667, 6668, 6669, 6998, 7000, 7001, 7005, 7014, 7071, 7077, 7080, 7288, 7401, 7443, 7474, 7493, 7537, 7547, 7548, 7634, 7657, 7777, 7779, 7911, 8112, 8123, 8125, 8126, 8139, 8161, 8200, 8291, 8333, 8334, 8377, 8378, 8443, 8500, 8545, 8554, 8649, 8686, 8800, 8834, 8880, 8883, 8888, 8889, 8983, 9000, 9001, 9002, 9003, 9009, 9010, 9042, 9051, 9080, 9090, 9100, 9151, 9191, 9200, 9295, 9333, 9418, 9443, 9527, 9530, 9595, 9653, 9700, 9711, 9869, 9944, 9981, 9999, 10000, 10001, 10162, 10243, 10333, 11001, 11211, 11300, 11310, 12300, 12345, 13579, 14000, 14147, 14265, 16010, 16030, 16992, 16993, 17000, 18001, 18081, 18245, 18246, 19999, 20000, 20547, 22105, 22222, 23023, 23424, 25000, 25105, 25565, 27015, 27017, 28017, 32400, 33338, 33890, 37215, 37777, 41795, 42873, 45554, 49151, 49152, 49153, 49154, 49155, 50000, 50050, 50070, 50100, 51106, 52869, 55442, 55553, 60001, 60010, 60030, 61613, 61616, 62078, 64738,
	}
	CheckTCP_port_range(ports)
}

func CheckTCP_port_range(ports []int) {
	//exit checker
	quit := false

	allowed_ports_number := 0

	mwg := maxedWaitGroup{
		current: make(chan int, 100),
		wg:      sync.WaitGroup{},
	}

	for _, port := range ports {

		//check the quit variable and terminate the checking.
		//have some bugs I may fix it in the future.
		if quit == true {
			log.Println("[*] Found more than 3 ports can access the Internet. Stop further testing.")
			return
		}

		mwg.Add()

		switch {
		case (quit == true):
			log.Println("[*] Found more than 3 ports can access the Internet. Stop further testing.")
			return
		default:
			go func(p int) {
				defer mwg.Done()
				time.Sleep(time.Second * time.Duration(rand.Intn(1)))
				resp := testHTTPEgress(p)
				if resp {
					allowed_ports_number++
				}
				return
			}(port)
		}
		if allowed_ports_number > 3 {
			quit = true
		}

	}
	// Wait for the work to complete
	mwg.Wait()

}

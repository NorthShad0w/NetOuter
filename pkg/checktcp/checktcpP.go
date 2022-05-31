package checktcp

import "strconv"

func CheckALLtcp() {
	for i := 1; i <= 65535; i++ {
		Checktcp("45.79.204.144", strconv.Itoa(i))
	}
}

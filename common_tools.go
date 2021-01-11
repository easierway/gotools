package gotools

import (
	"net"
	"strconv"
	"strings"
)

func stringToInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic("host address is wrong format! " + err.Error())
	}
	return v
}

// GetSeedForRandomCreation is to create the seed according to the host address
func GetSeedForRandomCreation() int64 {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic("can't get the host addresses." + err.Error())
	}

	for _, address := range addrs {

		// skip the loop address
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipSecs := strings.Split(ipnet.IP.To4().String(), ".")
				ret := int64(stringToInt(ipSecs[0])*255*255*255 + stringToInt(ipSecs[1])*255*255 +
					stringToInt(ipSecs[2])*255 + stringToInt(ipSecs[3]))
				return ret
			}
		}
	}
	return 0
}

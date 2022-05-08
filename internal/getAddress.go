package internal

import (
	"log"
	"net"
)

func GetInterfaceAddrs() (ipv4_addr map[string]string, ipv6_addr map[string]string) {
	itfs, err := net.Interfaces()
	if err != nil { // get interface
		log.Panicf(err.Error())
	}

	ipv4_addr = map[string]string{}
	ipv6_addr = map[string]string{}
	for _, itf := range itfs {
		addrs, err := itf.Addrs()
		if err != nil { // get addresses
			log.Panicf(err.Error())
		}
		for _, addr := range addrs {
			if ipnet := addr.(*net.IPNet).IP; ipnet.To4() != nil {
				ipv4_addr[itf.Name] = ipnet.To4().String()
			} else {
				ipv6_addr[itf.Name] = ipnet.To16().String()
			}
		}
	}
	return
}

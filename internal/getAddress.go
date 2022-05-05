package internal

import (
	"log"
	"net"
)

// Get IPv4 address from network interface. / 获取网络接口的IPv4地址。
// useful links / 参考链接:
// https://stackoverflow.com/questions/27410764/dial-with-a-specific-address-interface-golang
// https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6
func GetInterfaceIPv4Addr(interfaceName string) (addr string) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
		err      error
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		log.Panicf(err.Error())
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		log.Panicf(err.Error())
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		log.Panicf("Interface don't have IPv4 address. / 该接口无 IPv4 地址。")
	}
	return ipv4Addr.String()
}

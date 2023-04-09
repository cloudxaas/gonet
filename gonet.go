package cxnet

import (
	"net/netip"
	"net"
)


// NonLoopbackPrimaryIP returns the non loopback local IP of the host
// returns only 1 ip
func NonLoopbackPrimaryIP() netip.Addr {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return (netip.Addr{})
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return netip.MustParseAddr(ipnet.IP.String())
				//return ipnet.IP
			}
		}
	}
	return (netip.Addr{})
}

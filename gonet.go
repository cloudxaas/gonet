package cxnet

import (
	"net"
)


// NonLoopbackPrimaryIP returns the non loopback local IP of the host
// returns only 1 ip
func NonLoopbackPrimaryIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

package cxnet

import (
	"net"
)

func IPToBytes(ip net.IP) []byte {
    if ip4 := ip.To4(); ip4 != nil {
        return []byte(ip4)
    }
    return []byte(ip.To16())
}

// NonLoopbackPrimaryIP returns the non loopback local IP of the host
// returns only 1 ip
func NonLoopbackPrimaryIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

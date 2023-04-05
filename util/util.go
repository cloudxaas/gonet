package cxnetutil
import (
"net"
)
func IsPrivateSubnet(ipAddress net.IP) uint8 {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range PrivateRanges {
			// check if this ip is in a private range
			if InRange(r, ipAddress) == 1 {
				return 1
			}
		}
	}
	return 0
}

func InRange(r IpRange, ipAddress net.IP) uint8 {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return 1
	}
	return 0
}

var PrivateRanges = []IpRange{
	IpRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	/*
		IpRange{
			start: net.ParseIP("100.64.0.0"),
			end:   net.ParseIP("100.127.255.255"),
		},
	*/
	IpRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	/*
		IpRange{
			start: net.ParseIP("192.0.0.0"),
			end:   net.ParseIP("192.0.0.255"),
		},
	*/
	IpRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	/*
		IpRange{
			start: net.ParseIP("198.18.0.0"),
			end:   net.ParseIP("198.19.255.255"),
		},
	*/
}

//IpRange - a structure that holds the start and end of a range of ip addresses
type IpRange struct {
	start net.IP
	end   net.IP
}

func ListContainsIP(ipList []net.IP, ip net.IP) uint8 {
	for i := range ipList {
		if bytes.Equal(ipList[i], ip) {
			return 1
		}
	}
	return 0
}

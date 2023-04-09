package cxnetip
import (
	"net"
	"bytes"
)

func AppendSortedIPNetSlices(sorted *[]*net.IPNet, ipnet *net.IPNet) {
    if len(*sorted) == 0 {
        *sorted = append(*sorted, ipnet)
        return
    }

    var left, right, mid int
    left = 0
    right = len(*sorted) - 1

    for left <= right {
        mid = left + (right-left)/2
        if (*sorted)[mid].Contains(ipnet.IP) {
            return
        } else if bytes.Compare((*sorted)[mid].IP, ipnet.IP) < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    // Append the new IPNet to the end of the slice
    *sorted = append(*sorted, nil)

    // Shift the elements to the right of the insertion point one position to the right
    copy((*sorted)[left+1:], (*sorted)[left:len(*sorted)-1])

    // Insert the new IPNet at the insertion point
    (*sorted)[left] = ipnet
}

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

func ListContainsIP(ipList []netip.IPPrefix, ip netip.Addr) uint8 {
	for _, block := range ipList {
		if block.Contains(ip) {
			return 1
		}
	}
	return 0
}

func Listv4or6ContainsIP(ipListv4 []netip.IPPrefix, ipListv6 []netip.IPPrefix, ip netip.Addr) uint8 {
	if ip.Is4() {
		if ListContainsIP(ipListv4, ip) == 1 {
			return 1 // IPv4 and found in the list
		}
	} else {
		if ListContainsIP(ipListv6, ip) == 1 {
			return 2 // IPv6 and found in the list
		}
	}
	return 0 // Not found in either list
}

package cxnetip

import (
	"net/netip"
)

func AppendSortedNetIPPrefixSlices(sorted *[]netip.Prefix, ip netip.Prefix) {
        if len(*sorted) == 0 {
                (*sorted)[0] = ip
                *sorted = (*sorted)[:1] // Correct the slice length
                return
        }
        var left, right, mid int
        left = 0
        right = len(*sorted) - 1
        for left <= right {
                mid = left + (right-left)/2
                if (*sorted)[mid].Contains(ip.Addr()) {
                        return
                } else if (*sorted)[mid].Addr().Less(ip.Addr()) {
                        left = mid + 1
                } else {
                        right = mid - 1
                }
        }

        // Ensure we stay within the slice bounds
        *sorted = (*sorted)[:len(*sorted)+1] // Extend the slice length to accommodate the new element

        // Manually shift elements to avoid using append (which may allocate)
        copy((*sorted)[left+1:], (*sorted)[left:len(*sorted)-1])
        (*sorted)[left] = ip
}


func IsPrivateSubnet(ipAddress netip.Addr) uint8 {
	if ipAddress != (netip.Addr{}) {
		for _, r := range PrivateRanges {
			if InRange(r, ipAddress) == 1 {
				return 1
			}
		}
	}
	return 0
}

func InRange(r IpRange, ip netip.Addr) uint8 {
	if r.start.Compare(ip) <= 0 && r.end.Compare(ip) > 0 {
		return 1
	}
	return 0
}

var PrivateRanges = []netip.Prefix{
	netip.MustParsePrefix("10.0.0.0/8"),
	netip.MustParsePrefix("172.16.0.0/12"),
	netip.MustParsePrefix("192.168.0.0/16"),
}

func ListContainsIP(ipList []netip.Prefix, ip netip.Addr) uint8 {
	for _, block := range ipList {
		if block.Contains(ip) {
			return 1
		}
	}
	return 0
}

func Listv4or6ContainsIP(ipListv4, ipListv6 []netip.Prefix, ip netip.Addr) uint8 {
	if ip.Is4() {
		if ListContainsIP(ipListv4, ip) {
			return 1 // IPv4 and found in the list
		}
	} else if ip.Is6() {
		if ListContainsIP(ipListv6, ip) {
			return 2 // IPv6 and found in the list
		}
	}
	return 0 // Not found in either list
}

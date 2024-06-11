package cxnetip

import (
	"net/netip"
)

func ipAddrAsSlice(ip net.IP) [16]byte {
	var arr [16]byte
	copy(arr[:], ip.To16())
	return arr
}

func AppendSortedNetIPPrefixSlices(sorted *[]netip.Prefix, ip netip.Prefix) {
	if len(*sorted) == 0 {
		*sorted = append(*sorted, ip)
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

	// Append the new Prefix to the end of the slice
	*sorted = append(*sorted, netip.Prefix{})

	// Shift the elements to the right of the insertion point one position to the right
	copy((*sorted)[left+1:], (*sorted)[left:len(*sorted)-1])

	// Insert the new Prefix at the insertion point
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

var PrivateRanges = []IpRange{
	IpRange{
		start: netip.MustParseAddr("10.0.0.0"),
		end:   netip.MustParseAddr("10.255.255.255"),
	},
	IpRange{
		start: netip.MustParseAddr("172.16.0.0"),
		end:   netip.MustParseAddr("172.31.255.255"),
	},
	IpRange{
		start: netip.MustParseAddr("192.168.0.0"),
		end:   netip.MustParseAddr("192.168.255.255"),
	},
}

type IpRange struct {
	start netip.Addr
	end   netip.Addr
}

func ListContainsIP(ipList []netip.Prefix, ip netip.Addr) uint8 {
	for _, block := range ipList {
		if block.Contains(ip) {
			return 1
		}
	}
	return 0
}

func Listv4or6ContainsIP(ipListv4 []netip.Prefix, ipListv6 []netip.Prefix, ip netip.Addr) uint8 {
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

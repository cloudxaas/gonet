package cxnetip

import (
	"net/netip"
)

func AppendSortedNetIPSlices(sorted *[]netip.Addr, ip netip.Addr) {
	if len(*sorted) == 0 {
		*sorted = append(*sorted, ip)
		return
	}

	var left, right, mid int
	left = 0
	right = len(*sorted) - 1

	for left <= right {
		mid = left + (right-left)/2
		if (*sorted)[mid] == ip {
			return
		} else if (*sorted)[mid].Less(ip) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Append a zero-value Addr to the end of the slice
	*sorted = append(*sorted, (netip.Addr{}))

	// Shift the elements to the right of the insertion point one position to the right
	copy((*sorted)[left+1:], (*sorted)[left:len(*sorted)-1])

	// Insert the new Addr at the insertion point
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

func InRange(r IpRange, ipAddress netip.Addr) uint8 {
	if r.start.Compare(ipAddress) >= 0 && ipAddress.Compare(r.end) < 0  {
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

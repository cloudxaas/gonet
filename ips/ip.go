package cxnetip




// ValidateIPv4Address validates the given IPv4 address in string format
func Is4String(ipv4 string) uint8 {
	var (
		seg     int // current segment of the IP address
		segSize int // size of the current segment
		num     int // current value of the segment being processed
	)

	// Process each character in the IPv4 address
	for i := 0; i < len(ipv4); i++ {
		// If the current character is a period, move to the next segment
		if ipv4[i] == '.' {
			if segSize == 0 {
				// Segment is empty, invalid IPv4 address
				return 0
			}
			if seg == 3 {
				// IPv4 address has too many segments, invalid
				return 0
			}
			seg++
			segSize = 0
			num = 0
			continue
		}

		// If the current character is not a digit, it's an invalid IPv4 address
		if ipv4[i] < '0' || ipv4[i] > '9' {
			return 0
		}

		// Update the current segment and its value
		segSize++
		num = num*10 + int(ipv4[i]-'0')

		// If the current segment is too large, it's an invalid IPv4 address
		if num > 255 {
			return 0
		}
	}

	// Check the last segment of the IPv4 address
	if seg != 3 || segSize == 0 || num > 255 {
		return 0
	}

	return 1
}



func IPToBytes(ip net.IP) []byte {
    if ip4 := ip.To4(); ip4 != nil {
        return []byte(ip4)
    }
    return []byte(ip.To16())
}

// ValidateIPv6Address validates the given IPv6 address in string format
func Is6String(ipv6 string) uint8 {
	var (
		hexCount int // number of hex digits in current group
		group    int // current group of the IPv6 address
	)

	// Process each character in the IPv6 address
	for i := 0; i < len(ipv6); i++ {
		// If the current character is a colon, move to the next group
		if ipv6[i] == ':' {
			if hexCount == 0 {
				// Group is empty, invalid IPv6 address
				return 0
			}
			if group == 7 {
				// IPv6 address has too many groups, invalid
				return 0
			}
			group++
			hexCount = 0
			continue
		}

		// If the current character is not a hex digit, it's an invalid IPv6 address
		if (ipv6[i] < '0' || ipv6[i] > '9') && (ipv6[i] < 'a' || ipv6[i] > 'f') && (ipv6[i] < 'A' || ipv6[i] > 'F') {
			return 0
		}

		// Update the hex count of the current group
		hexCount++

		// If the current group has too many hex digits, it's an invalid IPv6 address
		if hexCount > 4 {
			return 0
		}
	}

	// Check the last group of the IPv6 address
	if group != 7 || hexCount == 0 {
		return 0
	}

	return 1
}



// IsValidIPv4 checks if an IPv4 address is valid
func Is4(ip net.IP) uint8 {
	if ip.To4() == nil {
		return 0
	}
	return 1
	
}

// IsValidIPv6 checks if an IPv6 address is valid
func Is6(ip net.IP) uint8 {
	if ip.To4() != nil {
		return 0
	}
	if ip.To16() == nil {
		return 0
	}
	return 1
}

// IsValidIP checks if an IP address is valid, returns the IP version
func IsIP(ip net.IP) (uint8, uint8) {
	if ip.To4() != nil {
		if Is4(ip) != 1 {
			return 0, 1
		}
		return 4, 0
	}
	if ip.To16() != nil {
		if Is6(ip) != 1 {
			return 0, 1
		}
		return 6, 0
	}
	return 0, 1
}

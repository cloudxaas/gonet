package cxnethttp

import (
    "unicode"
    "unicode/utf8"
	"net/netip"
"bytes"
)

//validate if domain name is valid including internationalized domain names
func IsDomainName(domain string) bool {
    if domain == "" || len(domain) > 255 {
        return false
    }
    if domain[len(domain)-1] == '.' {
        domain = domain[:len(domain)-1]
    }
    i, labelLen := 0, 0
    for i < len(domain) {
        if domain[i] == '.' {
            if labelLen < 1 || labelLen > 63 {
                return false
            }
            labelLen = 0
            i++
            continue
        }
        char, size := utf8.DecodeRuneInString(domain[i:])
        if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' {
            return false
        }
        i += size
        labelLen++
    }
    return labelLen >= 1 && labelLen <= 63
}


func IPFromXFF(header []byte) netip.Addr {
	if len(header) < 7 {
		return netip.Addr{}
	}

	// Split the header on commas and reverse the resulting slice
	addresses := bytes.Split(header, []byte(","))
	for i, j := 0, len(addresses)-1; i < j; i, j = i+1, j-1 {
		addresses[i], addresses[j] = addresses[j], addresses[i]
	}

	for _, addr := range addresses {
		// Remove any whitespace and double quotes from the address
		addr = bytes.TrimSpace(addr)
		addr = bytes.TrimPrefix(addr, []byte{'"'})
		addr = bytes.TrimSuffix(addr, []byte{'"'})

		// Check if the address is a valid IP
		ip, err := netip.ParseAddr(string(addr))
		if err == nil {
			return ip
		}
	}

	return netip.Addr{}
}

package ipvalidator

func ValidateIPv4Address(ipAddress []byte) bool {
    const maxOctets = 4
    octets := make([]int, maxOctets)
    numOctets := 0
    numDigits := 0
    
    for i := 0; i < len(ipAddress); i++ {
        c := ipAddress[i]
        if c == '.' {
            numOctets++
            if numOctets > maxOctets || numDigits == 0 || numDigits > 3 {
                return false
            }
            numDigits = 0
        } else if c >= '0' && c <= '9' {
            octets[numOctets] = octets[numOctets]*10 + int(c-'0')
            if octets[numOctets] > 255 {
                return false
            }
            numDigits++
        } else {
            return false
        }
    }
    
    return numOctets == maxOctets && numDigits > 0 && numDigits <= 3
}

func ValidateIPv6Address(ipAddress []byte) bool {
    const maxHextets = 8
    hextets := make([]uint16, maxHextets)
    numHextets := 0
    numDigits := 0
    
    for i := 0; i < len(ipAddress); i++ {
        c := ipAddress[i]
        if c == ':' {
            numHextets++
            if numHextets > maxHextets || numDigits == 0 || numDigits > 4 {
                return false
            }
            numDigits = 0
        } else if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
            hextets[numHextets] = hextets[numHextets]<<4 | uint16(c)
            numDigits++
        } else {
            return false
        }
    }
    
    return numHextets == maxHextets && numDigits > 0 && numDigits <= 4
}

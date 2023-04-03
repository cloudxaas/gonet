package cxnethttp

import (
    "unicode"
    "unicode/utf8"
)

//validate if domain name is valid including internationalized domain names
func IsValidDomainName(domain string) bool {
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

package cxnetmail

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	emailPattern = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

// ValidateEmail checks if the email is valid with zero allocation.
func ValidateEmail(email string) bool {
	return emailPattern.MatchString(email)
}

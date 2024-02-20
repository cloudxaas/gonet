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

// ParseEmailAddress parses a contact string into name and email with minimized allocations.
func ExtractNameAndEmail(contact string) (name, email string, err error) {
	// Use strings.IndexByte when searching for single characters to avoid allocations.
	start := strings.IndexByte(contact, '<')
	end := strings.IndexByte(contact, '>')

	if start == -1 || end == -1 || start >= end {
		err = fmt.Errorf("invalid contact format")
		return
	}

	// Avoid trimming and slicing strings unnecessarily by directly assigning.
	name = contact[:start]
	email = contact[start+1 : end]

	// Trim potential quotes around the name without allocating a new string.
	if len(name) > 0 && name[0] == '"' && name[len(name)-1] == '"' {
		name = name[1 : len(name)-1]
	}

	// Use strings.TrimSpace to minimize allocations instead of creating substrings.
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if !ValidateEmail(email) {
		err = fmt.Errorf("invalid email address")
		return
	}

	return
}

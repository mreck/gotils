package gotils

import (
	"strings"
)

func StringMatchesAny(s string, needles []string) bool {
	for _, needle := range needles {
		if s == needle {
			return true
		}
	}
	return false
}

func StringContainsAny(s string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(s, needle) {
			return true
		}
	}
	return false
}

func StringContainsAll(s string, needles []string) bool {
	for _, needle := range needles {
		if !strings.Contains(s, needle) {
			return false
		}
	}
	return true
}

func StringContainsNone(s string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(s, needle) {
			return false
		}
	}
	return true
}

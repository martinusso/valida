package validate

import (
	"regexp"
	"strings"
)

type validationFunc func(value interface{}) bool

type validation struct {
	fn   validationFunc
	size int
}

func Email(s string) bool {
	if Empty(s) {
		return false
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(s)
}

func Empty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func Max(value, length int) bool {
	return value <= length
}

func MaxLength(value string, length int) bool {
	return len(value) <= length
}

func Min(value, length int) bool {
	return value >= length
}

func MinLength(value string, length int) bool {
	return len(value) >= length
}

func Password(s string) bool {
	if len(s) < 6 {
		return false
	}

	hasChar, _ := regexp.MatchString("^.*[a-zA-Z]", s)
	if !hasChar {
		return false
	}

	hasNumber, _ := regexp.MatchString("^.*[0-9]", s)
	if !hasNumber {
		return false
	}

	hasSymbol, _ := regexp.MatchString(`^.*[!@#$%&*()_+=\-'"]`, s)
	if !hasSymbol {
		return false
	}
	return true
}

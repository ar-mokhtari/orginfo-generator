package validator

import "regexp"

func IsLuhnAlgorithm(input string, format *regexp.Regexp) bool {
	if format.MatchString(input) {
		return true
	}
	return false
}

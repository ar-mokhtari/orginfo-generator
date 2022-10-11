package convertor

import "strings"

func ToUpper(input string) (res string) {
	if input == "" {
		return
	}
	return strings.ToUpper(input[:1]) + input[1:]
}

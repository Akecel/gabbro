package utils

import (
	"strings"
	"unicode"
)

func CamelCaseToNormal(str string) string {
	var result string

	for i, r := range str {
		if unicode.IsUpper(r) && i != 0 {
			if i > 0 && str[i-1] != ' ' && !unicode.IsUpper(rune(str[i-1])) {
				result += " "
			}
			result += string(unicode.ToLower(r))
		} else {
			result += string(r)
		}
	}

	return result
}

func ContainsI(strA string, strB string) bool {
	return strings.Contains(
		strings.ToLower(strA),
		strings.ToLower(strB),
	)
}

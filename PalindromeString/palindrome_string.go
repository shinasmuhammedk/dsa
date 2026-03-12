package palindromestring

import (
	"strings"
	"unicode"
)

func PalindromeString(s string) bool {
	n := strings.ToLower(s)
	r := []rune{}

	for _, v := range n {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			r = append(r, v)
		}
	}

	i := 0
	j := len(r) - 1

	for i < j {
        if r[i] != r[j]{
            return false
        }
        i++
        j--    
	}
    return true
}

// 125. Valid Palindrome 
//

import (
	"strings"
	"unicode"

)

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
		left := rune(s[i])
		right := rune(s[j])

		if !unicode.IsLetter(left) && !unicode.IsDigit(left) {
			i++ 
			continue 
		}
		if !unicode.IsLetter(right) && !unicode.IsDigit(right) {
			j--
			continue 
		}
		if unicode.ToLower(left) != unicode.ToLower(right) {
			return false 
		}
		i++
		j--
	}
	return true 
}

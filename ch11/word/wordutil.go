package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune
	for _, c := range s {
		if unicode.IsLetter(c) {
			letters = append(letters, unicode.ToLower(c))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}

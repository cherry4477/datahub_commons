package common

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	UrlWordValidator        = regexp.MustCompile("[^a-zA-Z0-9_\\-]")
	UnicodeUrlWordValidator = regexp.MustCompile("[\\s]")
	EmailValidator          = regexp.MustCompile("[^a-zA-Z0-9_\\-\\.\\@]")
)

func init() {

}

func ValidateGeneralWord(word string) (string, bool) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return "", false
	}
	return word, true
}

// return validated word and if this input word is valid
func ValidateUrlWord(word string) (string, bool) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return "", false
	}
	return word, UrlWordValidator.FindString(word) == ""
}

func ValidateUnicodeUrlWord(word string) (string, bool) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return "", false
	}

	for _, r := range word {
		if r == utf8.RuneError {
			return word, false
		}
	}

	return word, UnicodeUrlWordValidator.FindString(word) == ""
}

func ValidateEmail(word string) (string, bool) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return "", false
	}
	index := strings.IndexByte(word, '@')
	if index <= 0 || index == len(word)-1 { // -1, 0 or len - 1
		return word, false
	}
	index = strings.IndexByte(word[index+1:], '@')
	if index > 0 {
		return word, false
	}
	return word, EmailValidator.FindString(word) == ""
}

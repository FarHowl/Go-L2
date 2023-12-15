package dev02

import (
	"errors"
	"strings"
	"unicode"
)

func UnpackString(str string) (string, error) {
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		if letter := rune(str[i]); unicode.IsLetter(letter) {
			if i+1 != len(str) && unicode.IsDigit(rune(str[i+1])) {
				builder.WriteRune(letter * rune(str[i+1]))
			} else {
				builder.WriteRune(letter)
			}
		} else {
			return "", errors.New("Incorrect string")
		}
	}

	return builder.String(), nil
}

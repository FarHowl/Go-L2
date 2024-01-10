package dev02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func writeToBuilder(builder *strings.Builder, i *int, str string, currentSymbol rune) error {
	var countBuilder strings.Builder
	for j := *i + 1; j < len(str); j++ {
		if digit := rune(str[j]); unicode.IsDigit(digit) {
			countBuilder.WriteRune(digit)
			*i++
		} else {
			break
		}
	}

	if countBuilder.String() != "" {
		letterCount, err := strconv.Atoi(countBuilder.String())
		if err != nil {
			return err
		}

		for i := 0; i < letterCount; i++ {
			builder.WriteRune(currentSymbol)
		}

	} else {
		builder.WriteRune(currentSymbol)
	}

	return nil
}

func UnpackString(str string) (string, error) {
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		if letter := rune(str[i]); unicode.IsLetter(letter) {
			writeToBuilder(&builder, &i, str, letter)
		} else if slash := rune(str[i]); slash == '/' {
			i++
			escapeSymbol := rune(str[i])
			writeToBuilder(&builder, &i, str, escapeSymbol)
		} else {
			return "", errors.New("incorrect string")
		}
	}

	return builder.String(), nil
}

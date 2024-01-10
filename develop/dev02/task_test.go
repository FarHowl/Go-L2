package dev02

import (
	"fmt"
	"testing"
)

func TestStringUnpack(t *testing.T) {
	stringUnpackTests := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
		"":         "",
		"qwe/4/5":  "qwe45",
		"qwe/45":   "qwe44444",
		"qwe//5":   "qwe/////",
	}

	str := "asd"
	str = "a" + str + "a"
	fmt.Println(str)

	for k, v := range stringUnpackTests {
		res, err := UnpackString(k)
		if err != nil {
			t.Error(err)
		} else {
			if res != v {
				t.Error("Not equal")
			} else {
				fmt.Println(k, " => ", v)
			}
		}
	}
}

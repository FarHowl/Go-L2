package dev04

import (
	"fmt"
	"testing"
)

func mapsAreEqual(map1, map2 map[string][]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || !stringsSlicesEqual(value1, value2) {
			fmt.Println(value1)
			fmt.Println(value2)
			return false
		}
	}

	return true
}

func stringsSlicesEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[string]int)
	for _, s := range str1 {
		count[s]++
	}

	for _, s := range str2 {
		count[s]--
		if count[s] < 0 {
			return false
		}
	}

	return true
}
func TestAnagrams(t *testing.T) {
	anagramsTests := map[string]map[string][]string{
		"пятак, пятка, тяпка, листок, слиток, столик, абоба": {
			"листок": {
				"слиток", "столик",
			},
			"пятак": {
				"тяпка", "пятка",
			},
		},
	}

	for input, refOutput := range anagramsTests {
		output := FindAnagram(&input)
		if !mapsAreEqual(output, refOutput) {
			t.Error("Not equal")
		}
	}
}

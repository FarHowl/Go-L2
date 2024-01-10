package dev04

import (
	"strings"
)

func FindAnagram(myArray *string) map[string][]string {
	setsMap := make(map[string]map[string]struct{})

	strSlice := strings.Split(*myArray, ", ")

	for _, valToCheck := range strSlice {
		valToCheck = strings.ToLower(valToCheck)

		isAlone := true
		for ref := range setsMap {
			if isAnagram(ref, valToCheck) {
				setsMap[ref][valToCheck] = struct{}{}
				isAlone = false
				break
			}
		}

		if isAlone {
			setsMap[valToCheck] = make(map[string]struct{})
		}
	}

	parsedMap := parseToMapMapSlice(setsMap)
	return parsedMap
}

func parseToMapMapSlice(setsMap map[string]map[string]struct{}) map[string][]string {
	newMap := make(map[string][]string, len(setsMap))

	for wordKey, mapStringStruct := range setsMap {
		for word := range mapStringStruct {
			newMap[wordKey] = append(newMap[wordKey], word)
		}
	}

	return newMap
}

func isAnagram(ref, valToCheck string) bool {
	refMap := make(map[rune]struct{})
	for _, v := range ref {
		refMap[v] = struct{}{}
	}

	for _, v := range valToCheck {
		if _, exists := refMap[v]; !exists {
			return false
		}
	}
	return true
}

package dev06

import (
	"reflect"
	"testing"
)

func TestParseFields(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{"1,2,3", []int{1, 2, 3}},
		{"1-3", []int{1, 2, 3}},
		{"1-3,5,7", []int{1, 2, 3, 5, 7}},
		{"", []int{}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := parseFields(testCase.input)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected: %v, Got: %v", testCase.expected, result)
			}
		})
	}
}

func TestSelectFields(t *testing.T) {
	testCases := []struct {
		fields         []string
		selectedFields []int
		expected       []string
	}{
		{[]string{"a", "b", "c", "d"}, []int{1, 3}, []string{"a", "c"}},
		{[]string{"1", "2", "3"}, []int{2}, []string{"2"}},
		{[]string{"foo", "bar", "baz"}, []int{4, 5}, []string{"", ""}},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			result := selectFields(testCase.fields, testCase.selectedFields)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected: %v, Got: %v", testCase.expected, result)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"abc", 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := parseInt(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected: %d, Got: %d", testCase.expected, result)
			}
		})
	}
}

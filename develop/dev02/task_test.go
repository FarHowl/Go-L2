package dev02

import "testing"

type stringUnpackTest struct {
	str string
}

var stringUnpackTests []stringUnpackTest {
	{"abcd"}
	{"a4bc5d"}
}

func TestStringUnpack(t *testing.T) {
	for _, test := range stringUnpackTests {
		UnpackString(test)
	}
}

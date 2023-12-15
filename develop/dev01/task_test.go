package dev01

import (
	"testing"
	"time"
)

type exactTimeTest struct {
	addr string
}

var exactTimeTests = []exactTimeTest{
	{"time.google.com"},
	{"time.windows.com"},
}

func TestGetExactTime(t *testing.T) {
	for _, test := range exactTimeTests {
		exactTime := GetExactTime(test.addr)
		if exactTime.Before(time.Now().Add(-3*time.Second)) || exactTime.After(time.Now().Add(3*time.Second)) {
			t.Errorf("Expected exact time to be close to the current time, got: %v", exactTime)
		}
	}
}

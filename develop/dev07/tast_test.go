package dev07

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		close(ch1)
	}()
	go func() {
		time.Sleep(5 * time.Second)
		close(ch2)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		close(ch3)
	}()

	orChannel := or(ch1, ch2, ch3)

	select {
	case <-orChannel:
	case <-time.After(6 * time.Second):
		t.Error("Timeout waiting for orChannel to close")
	}
}

func TestOrChannel(t *testing.T) {
	start := time.Now()
	OrChannel()
	duration := time.Since(start)

	if duration > 5*time.Second {
		t.Errorf("TestOrChannel took too long: %v", duration)
	}
}

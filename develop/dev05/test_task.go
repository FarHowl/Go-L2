package dev05

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := []byte("This is a test\nAnother line\nTest pattern line\nYet another line\n")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	oldStdout := os.Stdout
	oldArgs := os.Args

	os.Args = []string{"grep", "-i", "-n", "-A", "1", "pattern", tmpfile.Name()}

	var buf bytes.Buffer

	os.Stdout = oldStdout
	os.Args = oldArgs

	expected := "3:Test pattern line\n4:Yet another line\n"

	if buf.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}

func TestGrepCount(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	content := []byte("This is a test\nAnother line\nTest pattern line\nYet another line\n")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	oldStdout := os.Stdout
	oldArgs := os.Args

	os.Args = []string{"grep", "-c", "pattern", tmpfile.Name()}

	var buf bytes.Buffer

	os.Stdout = oldStdout
	os.Args = oldArgs

	expected := "1\n"

	if buf.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}

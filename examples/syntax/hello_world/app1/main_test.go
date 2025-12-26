package main

import (
	"bytes"
	"testing"
)

func TestDisplayHelloWorld(t *testing.T) {
	var buf bytes.Buffer
	DisplayHelloWorld(&buf)
	expected := "Hello, World!\n"
	got := buf.String()
	if expected != got {
		t.Errorf("Unexpected Output -> expected: '%v', got: '%v'", expected, got)
	}
}

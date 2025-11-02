package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// captureStdout replaces os.Stdout with a buffer and returns the captured output.
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // run the function that writes to stdout

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

func TestHello(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"hello", "World"}

	greeting := captureStdout(func() {
		main()
	})

	expected := "Hello, World!\n"
	if greeting != expected {
		t.Errorf("greeting = %s; want %s", greeting, expected)
	}
}

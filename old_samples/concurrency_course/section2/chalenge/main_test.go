package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	newMsg := "hello"

	wg.Add(1)
	go updateMessage(newMsg, &wg)
	wg.Wait()

	if newMsg != msg {
		t.Errorf("Message (%s) didn't update to: %s", msg, newMsg)
	}
}

func Test_printMessage(t *testing.T) {
	// spy on the stdout
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "hello"

	printMessage()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Errorf("Expected to find %s, but it is not there", msg)
	}
}

func Test_Main(t *testing.T) {
	// spy on the stdout
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") ||
		!strings.Contains(output, "Hello, cosmos!") ||
		!strings.Contains(output, "Hello, world!") {
		t.Error("Not found the right messages")
	}
}

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
	wg.Add(1)

	go updateMessage("hello!", &wg)
	wg.Wait()

	if msg != "hello!" {
		t.Errorf("Expected to [hello!], but got [%s]", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "hello world!"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if strings.Trim(output, "\n") != "hello world!" {
		t.Errorf("Expected to find [hello world!], but got %s", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find [Hello universe!], but got %s", output)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find [Hello cosmos!], but got %s", output)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find [Hello world!], but got %s", output)
	}
}

package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "$34320.00") {
		t.Error("wrong balance return.")
	}
}

// func Test_updateMessage(t *testing.T) {
// 	msg = "Hello, world!"

// 	wg.Add(1)
// 	go updateMessage("Hello, Golang!")
// 	wg.Wait()

// 	if msg != "Hello, Golang!" {
// 		t.Errorf("incorrect value in msg")
// 	}
// }

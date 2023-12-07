package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintLetter(t *testing.T) {
	stdOut := os.Stdout

	readerFileBytes, writterFileBytes, _ := os.Pipe()
	os.Stdout = writterFileBytes

	var wg sync.WaitGroup
	wg.Add(1)

	go printLetter("a", &wg)

	wg.Wait()

	_ = writterFileBytes.Close()

	result, _ := io.ReadAll(readerFileBytes)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "a") {
		t.Errorf("Expected to find 'a', but is not there")
	}
}

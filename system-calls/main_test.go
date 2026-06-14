package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestMainProgram(t *testing.T) {
	var stdoutBuf, stderrBuf bytes.Buffer
	config, err := NewCliConfig(WithOutStream(&stdoutBuf), WithErrorStream(&stderrBuf))
	if err != nil {
		t.Fatalf("Error creating config: %v", err)
	}
	app([]string{"ayoola", "golang", "error"}, config)
	output := stdoutBuf.String()
	fmt.Println("output:", output)

	if len(output) == 0 {
		t.Fatalf("Expected output, got nothing")
	}

	if !strings.Contains(output, "word ayoola is even") {
		t.Fatal("Expected output to contain 'word ayoola is even")
	}

	if !strings.Contains(output, "word golang is even") {
		t.Fatal("Expected output to contain 'word golang is even")
	}

	errors := stderrBuf.String()
	if len(errors) == 0 {
		t.Fatal("Expected errors, got nothing")
	}

	if !strings.Contains(errors, "word error is odd") {
		t.Fatal("Expected erros does not contain 'word error is odd'")
	}
}

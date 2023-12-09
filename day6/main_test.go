package main

import (
	"os"
	"testing"
)

func TestMultipleWaysOfWinning(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 288
	output := multiplyWaysOfWinning(string(input))
	if expectedOutput != output {
		t.Fatalf("expected output was %v and got %v", expectedOutput, output)
	}
}

func TestWaysOfWinningPart2(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 71503
	if output := waysOfWinningPart2(string(input)); output != expectedOutput {
		t.Fatalf("expected output was %v and got %v", expectedOutput, output)
	}

}

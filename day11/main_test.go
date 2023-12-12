package main

import (
	"os"
	"testing"
)

func TestSumOfLengths(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 374
	if output := SumOfLengths(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}
}

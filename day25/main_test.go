package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal("problem reading input file")
	}

	expectedOutput := 54
	if actualOutput := part1(string(input)); actualOutput != expectedOutput {
		t.Fatalf("Expected %v but got %v.", expectedOutput, actualOutput)
	}
}

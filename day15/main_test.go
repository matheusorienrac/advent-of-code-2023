package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal("Error reading file")
	}

	expectedOutput := 1320
	if actualOutput := part1(string(input)); actualOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, actualOutput)
	}

}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal("Error reading file")
	}

	expectedOutput := 145
	if actualOutput := part2(string(input)); actualOutput != expectedOutput {
		t.Errorf("Expected %d, got %d", expectedOutput, actualOutput)
	}

}

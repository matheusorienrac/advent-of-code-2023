package main

import (
	"os"
	"testing"
)

func TestTotalWinnings(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 6440
	if output := totalWinnings(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}
}

func TestTotalWinningsPart2(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 5905
	if output := totalWinningsPart2(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}
}

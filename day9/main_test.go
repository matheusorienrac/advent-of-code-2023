package main

import (
	"os"
	"testing"
)

func TestSumOfExtrapolatedValues(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 114
	if output := sumOfExtrapolatedValues(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}
}

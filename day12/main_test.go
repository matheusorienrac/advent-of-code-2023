package main

import (
	"os"
	"testing"
)

func TestSumOfArrangements(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 21
	if output := TestSumOfArrangements(string(input)); output != expectedOutput {
		t.Fatalf("Expected %s but got %s", expectedOutput, output)
	}

}

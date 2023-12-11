package main

import (
	"os"
	"testing"
)

func TestCountSteps(t *testing.T) {
	input, err := os.ReadFile("test.txt")

	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 2
	if output := countSteps(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}

}

func TestCountStepsPart2(t *testing.T) {
	input, err := os.ReadFile("test2.txt")

	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 6
	if output := countStepsPart2(string(input)); output != expectedOutput {
		t.Fatalf("Expected %v but got %v", expectedOutput, output)
	}

}

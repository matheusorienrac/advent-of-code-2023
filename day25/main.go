package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("problem reading input file")
	}

	fmt.Println(part1(string(input)))
}

func part1(input string) int {
	return 0

}

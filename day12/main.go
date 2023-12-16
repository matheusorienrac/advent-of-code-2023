package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(SumOfArrangements(input(string)))
}

func checkIfViableArrangement(springs string, report []int) bool {

}

func SumOfArrangements(input string) {
	rows := strings.Split(input, "\n")

	arrangements := make([]int, len(rows))
	for i := 0; i < len(rows); i++ {
		aux := strings.Split(rows[i], " ")
		springs := aux[0]
		report := strings.Split(aux[1], ",")

		// brute forcing the solution:
		var currSpringsCopy string
		copy(currSpringsCopy, springs)
		for j := 0; j < len(springs); j++ {
			if rows[i][j] == '?' {

			}
		}

	}

}

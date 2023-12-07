package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(countScratchCards(string(input)))
}

// removes entries that are actually just empty spaces
func trimSlice(xs []string) []string {
	for i := 0; i < len(xs); {
		if len(xs[i]) == 0 {
			xs = append(xs[:i], xs[i+1:]...)
		} else {
			i++
		}
	}
	return xs
}

func sumPoints(input string) int {
	rows := strings.Split(input, "\r\n")
	sum := 0
	for _, row := range rows {
		halfRow := strings.Split(row, "|")
		winnerNumbers := trimSlice(strings.Split(strings.Split(halfRow[0], ":")[1], " "))
		playedNumbers := trimSlice(strings.Split(halfRow[1], " "))

		sum += countPoints(winnerNumbers, playedNumbers)
	}
	return sum

}

// counts scratchcards for part2. Total rows at the end will be the number of scratchcards
func countScratchCards(input string) int {
	rows := strings.Split(input, "\r\n")
	// this is gonna make it easier for me to duplicate the right scratchCards when necessary
	scratchCards := map[int]string{}
	for i := 1; i <= len(rows); i++ {
		scratchCards[i] = rows[i-1]
	}
	for i := 0; i < len(rows); i++ {
		halfRow := strings.Split(rows[i], "|")
		gameNumberStr := trimSlice(strings.Split(strings.Split(halfRow[0], ":")[0], " "))[1]
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			log.Panic(err)
		}

		winnerNumbers := trimSlice(strings.Split(strings.Split(halfRow[0], ":")[1], " "))
		playedNumbers := trimSlice(strings.Split(halfRow[1], " "))

		rowsToAdd := countWinnerNumbers(winnerNumbers, playedNumbers)
		for j := 0; j < rowsToAdd; j++ {
			rows = append(rows, scratchCards[gameNumber+j+1])
		}

	}
	return len(rows)

}

// returns how many numbers in playedNumbers were winner numbers
func countWinnerNumbers(winnerNumbers []string, playedNumbers []string) int {
	winners := 0
	for _, playedNumber := range playedNumbers {
		for _, winnerNumber := range winnerNumbers {
			if winnerNumber == playedNumber {
				winners++
			}
		}
	}
	return winners

}

// returns how many points were scored in the provided set of playednumbers and winnernumbers
func countPoints(winnerNumbers []string, playedNumbers []string) int {
	points := 0
	for _, playedNumber := range playedNumbers {
		for _, winnerNumber := range winnerNumbers {
			if winnerNumber == playedNumber {
				if points > 0 {
					points *= 2
				} else {
					points = 1
				}
			}
		}
	}
	return points
}

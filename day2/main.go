package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic("couln't read input file.")
	}

	fmt.Println(sumOfPowers(string(input)))

}

// Part 2
func sumOfPowers(input string) int {
	// parse each row
	games := strings.Split(input, "\n")
	sum := 0
	for _, game := range games {
		fmt.Println(game)
		rounds := []map[string]int{}

		// we want to skip the blank space
		colonIndex := strings.Index(game, ":")

		currRound := map[string]int{}
		var currColor string
		var currQuantity string
		for i := colonIndex + 2; i < len(game); i++ {
			if unicode.IsDigit(rune(game[i])) {
				currQuantity += string(game[i])
			} else if unicode.IsLetter(rune(game[i])) {
				currColor += string(game[i])
			} else if game[i] == ',' {
				intQuantity, err := strconv.Atoi(currQuantity)
				if err != nil {
					log.Panic("CurrQuantity is not a valid number")
				}
				currRound[currColor] = intQuantity
				currColor = ""
				currQuantity = ""
			}
			if (game[i] == ';') || i == len(game)-1 {
				intQuantity, err := strconv.Atoi(currQuantity)
				if err != nil {
					log.Panic("CurrQuantity is not a valid number")
				}
				currRound[currColor] = intQuantity
				rounds = append(rounds, currRound)
				fmt.Println(currRound)
				currRound = map[string]int{}
				currColor = ""
				currQuantity = ""
			}

		}
		sum += getPower(rounds)
	}
	return sum
}

// returns the power for a game
func getPower(rounds []map[string]int) int {
	colors := []string{"red", "green", "blue"}
	minimumBag := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range rounds {
		for _, color := range colors {
			if round[color] > minimumBag[color] {
				minimumBag[color] = round[color]
			}
		}
	}
	power := 1
	for _, n := range minimumBag {
		power *= n

	}
	fmt.Println(fmt.Sprintf("MinimumBag: %v and Power: %v", minimumBag, power))
	return power

}

// Part 1
func isGamePossible(rounds []map[string]int, bag map[string]int) bool {
	colors := []string{"red", "green", "blue"}
	for _, round := range rounds {
		for _, color := range colors {
			if round[color] > bag[color] {
				return false
			}
		}
	}
	return true

}

func sumPossibleGames(input string, bag map[string]int) int {
	// parse each row
	games := strings.Split(input, "\n")
	sum := 0
	for gameID, game := range games {
		fmt.Println(game)
		rounds := []map[string]int{}

		// we want to skip the blank space
		colonIndex := strings.Index(game, ":")

		currRound := map[string]int{}
		var currColor string
		var currQuantity string
		for i := colonIndex + 2; i < len(game); i++ {
			if unicode.IsDigit(rune(game[i])) {
				currQuantity += string(game[i])
			} else if unicode.IsLetter(rune(game[i])) {
				currColor += string(game[i])
			} else if game[i] == ',' {
				intQuantity, err := strconv.Atoi(currQuantity)
				if err != nil {
					log.Panic("CurrQuantity is not a valid number")
				}
				currRound[currColor] = intQuantity
				currColor = ""
				currQuantity = ""
			}
			if (game[i] == ';') || i == len(game)-1 {
				intQuantity, err := strconv.Atoi(currQuantity)
				if err != nil {
					log.Panic("CurrQuantity is not a valid number")
				}
				currRound[currColor] = intQuantity
				rounds = append(rounds, currRound)
				fmt.Println(currRound)
				currRound = map[string]int{}
				currColor = ""
				currQuantity = ""
			}

		}
		if isGamePossible(rounds, bag) {
			sum += gameID + 1
		}
	}
	return sum
}

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
		log.Panic(err)
	}

	fmt.Println(waysOfWinningPart2(string(input)))
}

func waysOfWinningPart2(input string) int {
	time, distance := getDistanceAndTimePart2(input)

	waysOfWinning := 0
	for j := 0; j < time; j++ {
		// means it breaks the record.
		if j*(time-j) > distance {
			waysOfWinning++
		}
	}

	return waysOfWinning

}

// Part 2
func getDistanceAndTimePart2(input string) (time int, distance int) {
	rows := strings.Split(input, "\n")
	currNum := ""
	for i := 0; i < len(rows); i++ {
		for j, c := range rows[i] {
			if unicode.IsDigit(rune(c)) {
				currNum += string(c)
				// if we're at the last number of the row or next character is not a digit, means we can save currnumber
				if j == len(rows[i])-1 {
					intNum, err := strconv.Atoi(currNum)
					if err != nil {
						log.Panic(err)
					}
					// means we are parsing time
					if i == 0 {
						time = intNum
					} else {
						// means we are parsing distance
						distance = intNum
					}
					currNum = ""
				}
			}
		}
	}
	return
}

// ///////////// PART 1
func getDistanceAndTime(input string) (times []int, distances []int) {
	rows := strings.Split(input, "\n")
	currNum := ""
	for i := 0; i < len(rows); i++ {
		for j, c := range rows[i] {
			if unicode.IsDigit(rune(c)) {
				currNum += string(c)
				// if we're at the last number of the row or next character is not a digit, means we can save currnumber
				if j == len(rows[i])-1 || !unicode.IsDigit(rune(rows[i][j+1])) {
					intNum, err := strconv.Atoi(currNum)
					if err != nil {
						log.Panic(err)
					}
					// means we are parsing time
					if i == 0 {
						times = append(times, intNum)
					} else {
						// means we are parsing distance
						distances = append(distances, intNum)
					}

					currNum = ""
				}
			}
		}
	}
	return
}

func multiplyWaysOfWinning(input string) int {
	// parse the time
	times, distance := getDistanceAndTime(input)

	waysOfWinning := make([]int, len(times))
	fmt.Println(waysOfWinning)
	for i := 0; i < len(times); i++ {
		for j := 0; j < times[i]; j++ {
			// means it breaks the record.
			if j*(times[i]-j) > distance[i] {
				fmt.Println(times, distance, i, j)
				fmt.Println(j * (times[i] - j))
				waysOfWinning[i]++
			}
		}
	}

	mult := 1
	for _, count := range waysOfWinning {
		mult *= count
	}

	return mult

}

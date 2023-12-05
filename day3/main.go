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
		log.Panic("issue trygin to read the input txt file")
	}

	fmt.Println(sumPartNumbers2(string(input)))

}

type Number struct {
	strNum     *string
	row        *int
	startIndex *int
}

// goes through contestants list and if a gear, sum the multiplication of the elements
func sumOfGearRatios(gearContestants map[string][]string) int {
	sum := 0
	for _, contestant := range gearContestants {
		if len(contestant) == 2 {
			n1, err := strconv.Atoi(contestant[0])
			if err != nil {
				log.Panic("partNumber 1 is invalid")
			}

			n2, err := strconv.Atoi(contestant[1])
			if err != nil {
				log.Panic("partNumber 2 is invalid")
			}
			sum += n1 * n2
		}
	}
	return sum
}

func AddToGearContestants(symbol string, row int, col int, num string, gearContestants map[string][]string) {
	key := fmt.Sprintf("%sr%vc%v", symbol, row, col)
	if v := gearContestants[key]; len(v) > 0 {
		gearContestants[key] = append(gearContestants[key], num)
	} else {
		gearContestants[key] = []string{num}
	}
}

func isPartNumber(n Number, matrix []string, gearContestants map[string][]string) bool {
	// to the sides
	var partNumber bool
	if *n.startIndex > 0 {
		if c := rune(matrix[*n.row][*n.startIndex-1]); isSymbol(c) {
			AddToGearContestants(string(c), *n.row, *n.startIndex-1, *n.strNum, gearContestants)
			partNumber = true
		}
	}
	if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
		if c := rune(matrix[*n.row][*n.startIndex+len(*n.strNum)]); isSymbol(c) {
			AddToGearContestants(string(c), *n.row, *n.startIndex+len(*n.strNum), *n.strNum, gearContestants)
			partNumber = true
		}
	}

	// above
	if *n.row > 0 {
		if *n.startIndex > 0 {
			if c := rune(matrix[*n.row-1][*n.startIndex-1]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row-1, *n.startIndex-1, *n.strNum, gearContestants)
				partNumber = true
			}
		}
		for i := *n.startIndex; i < *n.startIndex+len(*n.strNum); i++ {
			if c := rune(matrix[*n.row-1][i]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row-1, i, *n.strNum, gearContestants)
				partNumber = true
			}
		}
		if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
			if c := rune(matrix[*n.row-1][*n.startIndex+len(*n.strNum)]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row-1, *n.startIndex+len(*n.strNum), *n.strNum, gearContestants)
				partNumber = true
			}
		}
	}
	// below
	if *n.row+1 < len(matrix) {
		if *n.startIndex > 0 {
			if c := rune(matrix[*n.row+1][*n.startIndex-1]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row+1, *n.startIndex-1, *n.strNum, gearContestants)
				partNumber = true
			}
		}
		for i := *n.startIndex; i < *n.startIndex+len(*n.strNum); i++ {
			if c := rune(matrix[*n.row+1][i]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row+1, i, *n.strNum, gearContestants)
				partNumber = true
			}
		}
		if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
			if c := rune(matrix[*n.row+1][*n.startIndex+len(*n.strNum)]); isSymbol(c) {
				AddToGearContestants(string(c), *n.row+1, *n.startIndex+len(*n.strNum), *n.strNum, gearContestants)
				partNumber = true
			}
		}
	}
	return partNumber
}

func pointerString(s string) *string {
	return &s
}
func pointerInt(i int) *int {
	return &i
}

// isSymbol doesnt seem to solve our problem here as it doesnt consider thing such as * as a symbol.
func isSymbol(r rune) bool {
	if !unicode.IsDigit(r) && r != rune('.') {
		return true
	}
	return false
}

// Part 1
func sumPartNumbers(input string) int {
	currNumber := Number{}
	sum := 0

	rows := strings.Split(input, "\r\n")
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if unicode.IsDigit(rune(rows[i][j])) {
				if currNumber.strNum == nil {
					currNumber.row = pointerInt(i)
					currNumber.startIndex = pointerInt(j)
					currNumber.strNum = pointerString("")
				}
				*currNumber.strNum += string(rows[i][j])
			} else {
				if currNumber.startIndex != nil {
					if isPartNumber(currNumber, rows, map[string][]string{}) {
						intNum, err := strconv.Atoi(*currNumber.strNum)
						if err != nil {
							log.Panic("string representation of the current number is invalid")
						}
						sum += intNum
					}
					currNumber = Number{}
				}
			}

		}

	}
	return sum
}

// Part 2
func sumPartNumbers2(input string) int {
	currNumber := Number{}
	gearContestants := make(map[string][]string)

	rows := strings.Split(input, "\r\n")
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if unicode.IsDigit(rune(rows[i][j])) {
				if currNumber.strNum == nil {
					currNumber.row = pointerInt(i)
					currNumber.startIndex = pointerInt(j)
					currNumber.strNum = pointerString("")
				}
				*currNumber.strNum += string(rows[i][j])
			} else {
				if currNumber.startIndex != nil {
					isPartNumber(currNumber, rows, gearContestants)
					currNumber = Number{}
				}
			}

		}

	}
	return sumOfGearRatios(gearContestants)
}

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

	fmt.Println(sumPartNumbers(string(input)))

}

type Number struct {
	strNum     *string
	row        *int
	startIndex *int
}

func isPartNumber(n Number, matrix []string) bool {
	// to the sides
	if *n.startIndex > 0 {
		if isSymbol(rune(matrix[*n.row][*n.startIndex-1])) {
			return true
		}
	}
	if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
		if isSymbol(rune(matrix[*n.row][*n.startIndex+len(*n.strNum)])) {
			return true
		}
	}

	// above
	if *n.row > 0 {
		if *n.startIndex > 0 {
			if isSymbol(rune(matrix[*n.row-1][*n.startIndex-1])) {
				return true
			}
		}
		for i := *n.startIndex; i < *n.startIndex+len(*n.strNum); i++ {
			if isSymbol(rune(matrix[*n.row-1][i])) {
				return true
			}
		}
		if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
			if isSymbol(rune(matrix[*n.row-1][*n.startIndex+len(*n.strNum)])) {
				return true
			}
		}
	}
	// below
	if *n.row+1 < len(matrix) {
		if *n.startIndex > 0 {
			if isSymbol(rune(matrix[*n.row+1][*n.startIndex-1])) {
				return true
			}
		}
		for i := *n.startIndex; i < *n.startIndex+len(*n.strNum); i++ {
			if isSymbol(rune(matrix[*n.row+1][i])) {
				return true
			}
		}
		if *n.startIndex+len(*n.strNum) < len(matrix[*n.row]) {
			if isSymbol(rune(matrix[*n.row+1][*n.startIndex+len(*n.strNum)])) {
				return true
			}
		}
	}
	return false
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

	rows := strings.Split(input, "\n")
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
					if isPartNumber(currNumber, rows) {
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

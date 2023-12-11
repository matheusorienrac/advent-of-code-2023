package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(sumOfExtrapolatedValuesPart2(string(input)))
}

// Part 2
func sumOfExtrapolatedValuesPart2(input string) int {
	rows := strings.Split(input, "\n")

	sum := 0
	for _, row := range rows {
		strNums := strings.Split(row, " ")
		nums := strNumsToIntNums(strNums)

		sequences := [][]int{nums}
		for lastSequence := nums; !allNumsEqual(lastSequence); {
			newSequence := []int{}
			for i := 0; i+1 < len(lastSequence); i++ {
				newSequence = append(newSequence, lastSequence[i+1]-lastSequence[i])
			}
			sequences = append(sequences, newSequence)
			lastSequence = newSequence
		}
		p := message.NewPrinter(language.English)
		p.Printf("%v %d\n", sequences, getExtrapolatedValuePart2(sequences))
		sum += getExtrapolatedValuePart2(sequences)
	}
	return sum
}

func getExtrapolatedValuePart2(sequences [][]int) int {
	extrapolatedValue := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		extrapolatedValue *= -1
		extrapolatedValue += sequences[i][0]
	}
	return extrapolatedValue
}

// PART 1
func strNumsToIntNums(strNums []string) []int {
	intNums := make([]int, len(strNums))
	for i, num := range strNums {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			log.Panic(err)
		}
		intNums[i] = intNum
	}

	return intNums
}

func allNumsEqual(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	firstNum := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != firstNum {
			return false
		}
	}
	return true
}

func getExtrapolatedValue(sequences [][]int) int {
	extrapolatedValue := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		extrapolatedValue += sequences[i][len(sequences[i])-1]
	}
	return extrapolatedValue
}

func sumOfExtrapolatedValues(input string) int {
	rows := strings.Split(input, "\n")

	sum := 0
	for i, row := range rows {
		strNums := strings.Split(row, " ")
		nums := strNumsToIntNums(strNums)

		sequences := [][]int{nums}
		for lastSequence := nums; !allNumsEqual(lastSequence); {
			newSequence := []int{}
			for i := 0; i+1 < len(lastSequence); i++ {
				newSequence = append(newSequence, lastSequence[i+1]-lastSequence[i])
			}
			sequences = append(sequences, newSequence)
			lastSequence = newSequence
		}
		p := message.NewPrinter(language.English)
		p.Printf("%v %d\n", sequences, getExtrapolatedValue(sequences))
		sum += getExtrapolatedValue(sequences)
		fmt.Println(i)
	}
	fmt.Println(len(rows))
	return sum
}

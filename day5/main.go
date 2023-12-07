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
	input, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(smallestLocation(string(input)))

}

func smallestLocation(input string) int64 {
	rows := strings.Split(input, "\n")

	// will always hold the result from the latest conversion. it ofc begins as seed
	currNumsStr := strings.Split(strings.TrimSpace(strings.Split(rows[0], ":")[1]), " ")

	currNums := make([]int64, len(currNumsStr))
	for i := 0; i < len(currNumsStr); i++ {
		iNum, err := strconv.ParseInt(currNumsStr[i], 10, 64)
		if err != nil {
			log.Panic(err)
		}
		currNums[i] = iNum
	}

	newNums := make([]int64, len(currNums))
	// starts at 1 cause we have already parsed the seed
	for _, row := range rows {
		// means it is either empty or a \n
		if len(row) <= 1 {
			continue
		} else if !unicode.IsDigit(rune(row[0])) {
			// means it is the name of the map and we finished the previous one.
			fmt.Println("____________________")
			continue
		}

		// parse row into numbers
		splitRow := strings.Split(row, " ")
		target, err := strconv.ParseInt(splitRow[0], 10, 64)
		if err != nil {
			log.Panic(err)
		}
		origin, err := strconv.ParseInt(splitRow[1], 10, 64)
		if err != nil {
			log.Panic(err)
		}
		rng, err := strconv.ParseInt(splitRow[2], 10, 64)
		if err != nil {
			log.Panic(err)
		}

		for i := 0; i < len(currNums){
			fmt.Printf(fmt.Sprintf("%v corresponde a ", currNums[i]))
			if currNums[i] >= origin && currNums[i] < origin+rng {
				newNums = append(newNums, currNums[i]-origin+target)
				currNums = append(currNums[:i], currNums[i+1:])
				fmt.Printf(fmt.Sprintf("%v\n", newNums[i]))
			} else {
				i++
				fmt.Printf(fmt.Sprintf("%v\n", currNums[i]))
			}
		}
		if

	}
	var lowest int64
	lowest = currNums[0]
	for i := 0; i < len(currNums); i++ {
		if currNums[i] < lowest {
			lowest = currNums[i]
		}
	}
	return lowest
}

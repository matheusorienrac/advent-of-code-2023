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

	// will put the translated numbers here
	newNums := make([]int64, len(currNums))
	copy(newNums, currNums)

	// starts at 1 cause we have already parsed the seed
	for _, row := range rows {
		// means it is either empty or a \n
		if len(row) <= 1 {
			continue
		} else if !unicode.IsDigit(rune(row[0])) {
			// means it is the name of the map and we finished the previous one.
			fmt.Println(row)
			fmt.Println("______________________")
			copy(currNums, newNums)
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

		fmt.Println(target, origin, rng)

		for i := 0; i < len(currNums); i++ {
			if currNums[i] >= origin && currNums[i] < origin+rng {
				newNums[i] = currNums[i] - origin + target
			}
		}
	}
	var lowest int64
	lowest = newNums[0]
	for i := 0; i < len(newNums); i++ {
		if newNums[i] < lowest {
			lowest = newNums[i]
		}
	}
	return lowest
}

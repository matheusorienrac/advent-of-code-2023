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

	fmt.Println(smallestLocation2(string(input)))

}

////////////////////////////////////////////////////////////////////////////
/// PART 2
////////////////////////////////////////////////////////////////////////////

func smallestLocation2(input string) int64 {
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
	newNums := []int64{}

	// starts at 1 cause we have already parsed the seed
	for rowIndex, row := range rows {
		// means it is either empty or a \n
		if len(row) <= 1 {
			fmt.Println("________________")
			continue
		} else if !unicode.IsDigit(rune(row[0])) {
			// means it is the name of the map and we finished the previous one.
			fmt.Println(row)
			// we only want to do this after we finish the first mapping
			if rowIndex > 2 {

				// append nums that were not in the previously processed map
				newNums = append(newNums, currNums...)
				currNums = make([]int64, len(newNums))
				copy(currNums, newNums)
				// empty it again
				newNums = []int64{}

			}
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

		// iterate by two so we skip the ranges and only get the start of the sequence
		for i := 0; i < len(currNums); {
			// checks if original number in the sequence is in the map range
			if currNums[i] >= origin && currNums[i] < origin+rng {
				newNums = append(newNums, currNums[i]-origin+target)
				// if the seed number + range are not completely inside the map origin + rng it means we need to break into two pairs
				if currNums[i]+currNums[i+1] > origin+rng {
					newNums = append(newNums, currNums[i+1]-((currNums[i]+currNums[i+1])-(origin+rng)))

					// This part exceeded the origin + rng and therefore still needs to be mapped
					currNums = append(currNums, origin+rng)
					currNums = append(currNums, currNums[i]+currNums[i+1]-(origin+rng))
				} else {
					newNums = append(newNums, currNums[i+1])
				}
				currNums = append(currNums[:i], currNums[i+2:]...)
			} else {
				//  original is not in the map range, but the sequence is
				if currNums[i] < origin && currNums[i]+currNums[i+1] <= origin+rng && currNums[i]+currNums[i+1] > origin {
					// This part is below the origin and therefore will still need to be mapped
					currNums = append(currNums, currNums[i])
					currNums = append(currNums, origin-currNums[i])

					// This part was processed and will move on to the newNumbers slice.
					newNums = append(newNums, target)
					newNums = append(newNums, currNums[i]+currNums[i+1]-origin)

					// delete already processed number from currNumber slice
					currNums = append(currNums[:i], currNums[i+2:]...)

				} else if currNums[i] < origin && currNums[i]+currNums[i+1] > origin+rng { // means end of sequence is not inside the map
					// This part is below the origin and therefore will still need to be mapped
					currNums = append(currNums, currNums[i])
					currNums = append(currNums, origin-currNums[i])

					// This was processed
					newNums = append(newNums, target)
					newNums = append(newNums, rng)

					// This still has to be processed
					currNums = append(currNums, origin+rng)
					currNums = append(currNums, currNums[i]+currNums[i+1]-(origin+rng))

					// delete already processed number from currNumber slice
					currNums = append(currNums[:i], currNums[i+2:]...)
				} else {
					// means we didnt delete anything from the currNums slice and therefore we can add 2 to the index
					i += 2
				}

			}
			fmt.Println("currNums", currNums)
			fmt.Println("NewNums", newNums)

		}
		fmt.Println(newNums)
	}
	var lowest int64
	// append the currNums that were not in the last processed map
	newNums = append(newNums, currNums...)
	lowest = newNums[0]
	// at the end, the ranges wont matter because the smallest will always be the first number of the sequence
	for i := 0; i < len(newNums); i += 2 {
		if newNums[i] < lowest {
			lowest = newNums[i]
		}
	}
	return lowest
}

// //////////////////////////////////////////////////////////////////////////
// / PART 1
// //////////////////////////////////////////////////////////////////////////
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
				fmt.Println(fmt.Sprintf("%v corresponds to ", currNums[i]))
				newNums[i] = currNums[i] - origin + target
				fmt.Println(fmt.Sprintf("%v\n", newNums[i]))

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

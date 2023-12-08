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
	newNums := make([]int64, len(currNums))
	copy(newNums, currNums)

	// starts at 1 cause we have already parsed the seed
	for _, row := range rows {
		// means it is either empty or a \n
		if len(row) <= 1 {
			fmt.Println("________________")
			continue
		} else if !unicode.IsDigit(rune(row[0])) {
			// means it is the name of the map and we finished the previous one.
			fmt.Println(row)
			currNums = make([]int64, len(newNums))
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

		// skip the ranges
		for i := 0; i < len(currNums); i += 2 {
			// checks if original number in the sequence is in the map range
			if currNums[i] >= origin && currNums[i] < origin+rng {
				newNums[i] = currNums[i] - origin + target
				// if the seed number + range are not completely inside the map origin + rng it means we need to break into two pairs
				if currNums[i]+currNums[i+1] > origin+rng {
					newNums[i+1] = currNums[i+1] - ((currNums[i] + currNums[i+1]) - (origin + rng))
					// adding the start of the new pair
					newNums = append(newNums, origin+rng)
					// adding the range of the new pair
					newNums = append(newNums, currNums[i]+currNums[i+1]-(origin+rng))
				}
			} else {
				// means original not in map range
				if currNums[i] < origin {
					//  original is not in the map range, but the sequence is
					if currNums[i]+currNums[i+1] <= origin+rng && currNums[i]+currNums[i+1] > origin {
						newNums[i+1] = origin - currNums[i]

						// adding the start of the new pair.
						newNums = append(newNums, target)
						// adding the range of the new pair
						newNums = append(newNums, currNums[i]+currNums[i+1]-origin)

					} else if currNums[i]+currNums[i+1] > origin+rng { // means end of sequence is not inside the map
						newNums[i+1] = origin - currNums[i]

						// adding the start of the new pair
						newNums = append(newNums, target)
						// adding the range of the new pair
						newNums = append(newNums, rng)

						// adding the start of the new pair
						newNums = append(newNums, origin+rng)
						// adding the range of the new pair
						newNums = append(newNums, currNums[i]+currNums[i+1]-(origin+rng))
					}
				}
			}
		}
		fmt.Println(newNums)
	}
	var lowest int64
	fmt.Println(len(newNums))
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

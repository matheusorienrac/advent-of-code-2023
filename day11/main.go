package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(SumOfLengths(string(input)))
}

func expandUniverse(input string) []string {
	rows := strings.Split(input, "\n")
	// gonna store all i's and j's that have galaxies
	// will hold the values of J's that cannot be expanded
	for i := 0; i < len(rows); i++ {
		expandCurrRow := true
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == '#' {
				expandCurrRow = false
				break
			}
		}
		if expandCurrRow {
			rows = append(rows[:i+1], rows[i:]...)
			// so we skip the newly added row
			i++
		}
	}

	for j := 0; j < len(rows[0]); j++ {
		expandCurrCol := true
		for i := 0; i < len(rows); i++ {
			if rows[i][j] == '#' {
				expandCurrCol = false
				break
			}
		}
		if expandCurrCol {
			for i := 0; i < len(rows); i++ {
				fmt.Println("neste j:", j, rows[i])
				fmt.Println("antes", len(rows[i]))
				rows[i] = rows[i][:j+1] + rows[i][j:]
				fmt.Println("depois", len(rows[i]))
			}
			// skip the newly added col
			j++
		}
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	return rows

}

func SumOfLengths(input string) int {
	galaxyLocations := [][]int{}
	rows := expandUniverse(input)

	// grabbing the locations and saving them
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == '#' {
				galaxyLocations = append(galaxyLocations, []int{i, j})
			}
		}
	}

	lengths := make([]int, len(galaxyLocations)*(len(galaxyLocations)-1)/2)
	// for each Galaxy location, figure otu the distance to the other locations
	for i := 0; i < len(galaxyLocations)-1; i++ {
		for j := i + 1; j < len(galaxyLocations); j++ {
			distance := math.Abs(float64(galaxyLocations[j][0]-galaxyLocations[i][0])) + math.Abs(float64(galaxyLocations[j][1]-galaxyLocations[i][1]))
			lengths = append(lengths, int(distance))
		}
	}

	sum := 0
	for _, length := range lengths {
		sum += length
	}
	return sum
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(countStepsPart2(string(input)))

}

func countStepsPart2(input string) int {
	rows := strings.Split(input, "\n")

	stepsTaken := 0
	directions := rows[0]

	nodes := map[string][]string{}
	nodesThatEndInA := []string{}
	// putting the directions into a map so its a bit more convenient to go through them for long cases
	// starting at 2 cause row 0 is directions and row 1 is empty space.
	for i := 2; i < len(rows); i++ {
		row := strings.ReplaceAll(rows[i], " ", "")
		row = strings.ReplaceAll(row, "(", "")
		row = strings.ReplaceAll(row, ")", "")
		splitRow := strings.Split(row, "=")
		node := splitRow[0]
		paths := strings.Split(splitRow[1], ",")
		nodes[node] = paths
		if node[2] == 'A' {
			nodesThatEndInA = append(nodesThatEndInA, node)
		}
	}

	// now we'll count the steps
	for currNodes := nodesThatEndInA; !allNodesEndWithZ(currNodes); {
		for i := 0; i < len(directions) && !allNodesEndWithZ(currNodes); i++ {
			for j := 0; j < len(currNodes); j++ {
				fmt.Println(currNodes, i, stepsTaken)
				if directions[i] == 'L' {
					currNodes[j] = nodes[currNodes[j]][0]
				} else {
					currNodes[j] = nodes[currNodes[j]][1]
				}
			}
			stepsTaken++
		}
	}

	return stepsTaken
}

func allNodesEndWithZ(nodes []string) bool {
	for i := 0; i < len(nodes); i++ {
		if nodes[i][2] != 'Z' {
			return false
		}
	}
	return true
}

// PART 1
func countSteps(input string) int {
	rows := strings.Split(input, "\n")

	stepsTaken := 0
	directions := rows[0]

	nodes := map[string][]string{}
	// putting the directions into a map so its a bit more convenient to go through them for long cases
	// starting at 2 cause row 0 is directions and row 1 is empty space.
	for i := 2; i < len(rows); i++ {
		row := strings.ReplaceAll(rows[i], " ", "")
		row = strings.ReplaceAll(row, "(", "")
		row = strings.ReplaceAll(row, ")", "")
		splitRow := strings.Split(row, "=")
		paths := strings.Split(splitRow[1], ",")
		nodes[splitRow[0]] = paths
	}

	// now we'll count the steps
	for currNode := "AAA"; currNode != "ZZZ"; {
		for i := 0; i < len(directions) && currNode != "ZZZ"; i++ {
			if directions[i] == 'L' {
				currNode = nodes[currNode][0]
			} else {
				currNode = nodes[currNode][1]
			}
			stepsTaken++
		}
	}

	return stepsTaken
}

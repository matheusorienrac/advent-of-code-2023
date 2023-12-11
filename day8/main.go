package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(countStepsPart2(string(input)))

}

type StepsToZ struct {
	startingNode string
	stepsTaken   int
}

// takes too freaking long to run, will try the parallel route
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

	var wg sync.WaitGroup
	fmt.Println(nodesThatEndInA)
	time.Sleep(time.Second)
	// if we find out the common steps taken for all starting nodes, we'll have our answer. I'll start a goroutine for each that will save how many steps it took for them to reach a node ending with a Z into this list

	c := make(chan StepsToZ, 50)

	// now we'll count the steps
	keepGoing := true
	for _, startingNode := range nodesThatEndInA {
		wg.Add(1)
		go findNodesThatEndInAZ(startingNode, directions, nodes, &keepGoing, c, &wg)
	}

	wg.Add(1)
	go checkIfCommonSteps(nodesThatEndInA, c, &keepGoing, &wg)

	wg.Wait()

	return stepsTaken
}

func findNodesThatEndInAZ(startingNode string, directions string, nodes map[string][]string, keepGoing *bool, c chan StepsToZ, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Println("desligaram o findnodes")
	}()
	stepsTaken := 0
	for currNode := startingNode; *keepGoing; {
		for i := 0; i < len(directions); i++ {
			if directions[i] == 'L' {
				currNode = nodes[currNode][0]
			} else {
				currNode = nodes[currNode][1]
			}
			stepsTaken++
			if currNode[2] == 'Z' {
				// the answer is greater than 1 trillion. This is just so things run a bit faster
				if stepsTaken >= 1000000000000 {
					c <- StepsToZ{startingNode: startingNode, stepsTaken: stepsTaken}
				}
			}
		}
	}

}

func checkIfCommonSteps(nodesThatEndInA []string, c chan StepsToZ, keepGoing *bool, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("desligaram o check if comon steps")
		wg.Done()
	}()
	// receive from channel
	stepsTakenEachStartingNode := map[string][]int{}
	for *keepGoing {
		for i := 0; i < len(c) && i < 50; i++ {
			elem := <-c
			stepsTakenEachStartingNode[elem.startingNode] = append(stepsTakenEachStartingNode[elem.startingNode], elem.stepsTaken)
		}

		// we need to find a stepsTaken number that is exactly the same for all 6 starting numbers.
		matches := 0
		deleteCurrent := false
		for z := 0; z < len(stepsTakenEachStartingNode[nodesThatEndInA[0]]); {
			stepsTaken := stepsTakenEachStartingNode[nodesThatEndInA[0]][z]
			for i := 1; i < len(nodesThatEndInA); i++ {
				for j := 0; j < len(stepsTakenEachStartingNode[nodesThatEndInA[i]]); j++ {
					if stepsTaken == stepsTakenEachStartingNode[nodesThatEndInA[i]][j] {
						fmt.Println("we have a match", stepsTaken, stepsTakenEachStartingNode[nodesThatEndInA[i]][j], nodesThatEndInA[i])
						matches++
						break
					} else if stepsTaken < stepsTakenEachStartingNode[nodesThatEndInA[i]][j] {
						// also delete all elements in nodesThatEndInA[i] up to this point, cause they cant be the solution either
						stepsTakenEachStartingNode[nodesThatEndInA[i]] = stepsTakenEachStartingNode[nodesThatEndInA[i]][j:]
						deleteCurrent = true
						break
					}
				}
				// means it broke out of the previous loop without finding any matches and the last number it tried to compare itself to is higher than it, which means it cant be the solution.
				if deleteCurrent {
					// delete from the list so we dont check this one again
					if len(stepsTakenEachStartingNode[nodesThatEndInA[0]]) == 1 {
						stepsTakenEachStartingNode[nodesThatEndInA[0]] = []int{}
					} else {
						stepsTakenEachStartingNode[nodesThatEndInA[0]] = stepsTakenEachStartingNode[nodesThatEndInA[0]][z+1:]
					}
					break
				}
			}
			if matches == len(nodesThatEndInA)-1 {
				// so all the other go routines stop running
				*keepGoing = false
				fmt.Println(stepsTakenEachStartingNode[nodesThatEndInA[0]][z])
				return
			} else {
				// reset everything and keep looking
				// means we deleted a number so we dont need to iterate z cause the list shrinked
				if deleteCurrent {
					deleteCurrent = false
				} else {
					z++
				}
				matches = 0
			}
		}
	}
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

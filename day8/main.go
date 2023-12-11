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
	defer wg.Done()
	fmt.Println("startei o node", startingNode)
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
				// the answer is greater than 100 million. This is just so things run a bit faster
				if stepsTaken >= 100000000 {
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
	stepsTakenEachStartingNode := map[string]map[int]bool{}
	for *keepGoing {
		for i := 0; i < len(c) && i < 50; i++ {
			elem := <-c
			if stepsTakenEachStartingNode[elem.startingNode] == nil {
				stepsTakenEachStartingNode[elem.startingNode] = map[int]bool{}
			}
			stepsTakenEachStartingNode[elem.startingNode][elem.stepsTaken] = true
		}

		// we need to find a stepsTaken number that is exactly the same for all 6 starting numbers.
		matches := 0
		for stepsTaken, _ := range stepsTakenEachStartingNode[nodesThatEndInA[0]] {
			fmt.Println(stepsTaken)
			for i := 1; i < len(nodesThatEndInA); i++ {
				if _, ok := stepsTakenEachStartingNode[nodesThatEndInA[i]][stepsTaken]; ok {
					matches++
				} else {
					break
				}
			}
			if matches == len(nodesThatEndInA)-1 {
				// so all the other go routines stop running
				*keepGoing = false
				fmt.Println(stepsTaken)
				return
			} else {
				// reset everything and keep looking
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

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

// safe for concurrency
type SafeStepsTakenEachStartingNode struct {
	mu       sync.Mutex
	StepsMap map[string][]int
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
	stepsTakenEachStartingNode := SafeStepsTakenEachStartingNode{
		StepsMap: map[string][]int{},
	}
	// now we'll count the steps
	keepGoing := true
	for _, startingNode := range nodesThatEndInA {
		wg.Add(1)
		go func() {
			defer wg.Done()
			findNodesThatEndInAZ(startingNode, directions, nodes, &keepGoing, stepsTakenEachStartingNode)
		}()
	}

	for keepGoing {
		time.Sleep(time.Second * 5)
		stepsTaken = checkIfCommonSteps(nodesThatEndInA, stepsTakenEachStartingNode, &keepGoing)
	}

	wg.Wait()

	return stepsTaken
}

func findNodesThatEndInAZ(startingNode string, directions string, nodes map[string][]string, keepGoing *bool, stepsTakenEachStartingNode SafeStepsTakenEachStartingNode) {
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
				stepsTakenEachStartingNode.mu.Lock()
				if stepsTakenEachStartingNode.StepsMap[startingNode] == nil {
					stepsTakenEachStartingNode.StepsMap[startingNode] = map[int]bool{}
				}
				stepsTakenEachStartingNode.StepsMap[startingNode][stepsTaken] = true
				stepsTakenEachStartingNode.mu.Unlock()
			}
		}
	}
}

func checkIfCommonSteps(nodesThatEndInA []string, stepsTakenEachStartingNode SafeStepsTakenEachStartingNode, keepGoing *bool) int {
	// after 30 seconds, check if there is a common number of steps taken for all starting nodes, and that is our solution
	time.Sleep(2 * time.Second)
	stepsTakenEachStartingNode.mu.Lock()

	// we need to find a stepsTaken number that is exactly the same for all 6 starting numbers.
	matches := 0
	for stepsTaken, _ := range stepsTakenEachStartingNode.StepsMap[nodesThatEndInA[0]] {
		for i := 1; i < len(nodesThatEndInA); i++ {
			if _, ok := stepsTakenEachStartingNode.StepsMap[nodesThatEndInA[i]][stepsTaken]; ok {
				matches++
			}
		}
		if matches == len(nodesThatEndInA)-1 {
			// means we found our solution
			fmt.Println(stepsTaken)
			// so all the other go routines stop running
			*keepGoing = false
			return stepsTaken
		} else {
			// reset everything and keep looking
			matches = 0
		}
	}
	stepsTakenEachStartingNode.mu.Unlock()
	return -1

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

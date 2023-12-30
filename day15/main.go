package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part2(string(input)))
}

type contents struct {
	Label     string
	FocalLens int
}

func LabelIndexInBox(label string, boxIndex int, boxes [][]contents) (int, error) {
	for i, c := range boxes[boxIndex] {
		if c.Label == label {
			return i, nil
		}
	}
	return -1, errors.New("Label not found in box")
}

func part2(input string) int {
	rows := strings.Split(strings.TrimSpace(input), ",")
	boxes := make([][]contents, 256, 256)
	for _, row := range rows {
		label := ""
		for _, c := range row {
			switch c {
			case '-':
				if index, err := LabelIndexInBox(label, hash(label), boxes); err == nil {
					boxes[hash(label)] = append(boxes[hash(label)][:index], boxes[hash(label)][index+1:]...)
				}
			case '=':
				focalLens := strings.Split(row, "=")[1]
				focalLensInt, err := strconv.Atoi(focalLens)
				if err != nil {
					log.Fatal(err)
				}
				if index, err := LabelIndexInBox(label, hash(label), boxes); err == nil {
					boxes[hash(label)][index].FocalLens = focalLensInt
				} else {
					boxes[hash(label)] = append(boxes[hash(label)], contents{label, focalLensInt})
				}
			default:
				label += string(c)
			}
		}
	}

	sum := 0
	for boxIndex, box := range boxes {
		for labelIndex, c := range box {
			fmt.Println(boxIndex, labelIndex, c, (boxIndex+1)*(labelIndex+1)*c.FocalLens)
			sum += (boxIndex + 1) * (labelIndex + 1) * c.FocalLens
		}
	}

	return sum
}

func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
		sum *= 17
		sum = sum % 256
	}
	fmt.Println(s, sum)
	return sum
}

func part1(input string) int {
	rows := strings.Split(strings.TrimSpace(input), ",")

	sum := 0
	for _, row := range rows {
		sum += hash(row)
	}

	return sum
}

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err, "cant read file for some reason")
	}
	fmt.Println(firstAndLastPart2(string((input))))
}

// if the provided string starts with a number in text, returns it.
func numberInStartOfString(s string) (*string, error) {
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, word := range numbers {
		if len(word) <= len(s) {
			for i := 0; i < len(s); i++ {
				if s[i] != word[i] {
					break
				} else if i == len(word)-1 {
					return &word, nil
				}
			}
		}

	}
	return nil, errors.New("no number in start of string")
}

func firstAndLastPart2(input string) int {
	numberMappings := map[string]byte{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	sum := 0
	var first, last *byte
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' || i == len(input)-1 {
			if first != nil && last != nil {
				currNumber, err := strconv.Atoi(fmt.Sprint("", string(*first), string(*last)))
				fmt.Println(currNumber)
				if err != nil {
					fmt.Println(string(*first))
					fmt.Println(string(*last))
					log.Panic("first and last number are not valid")
				}
				sum += currNumber
				first, last = nil, nil

			}
			continue

		}
		if unicode.IsDigit(rune(input[i])) {
			if first == nil {
				n := input[i]
				first = &n
				last = first
			} else {
				n := input[i]
				last = &n
			}
		} else {
			bounds := len(input)
			if i+5 < len(input) {
				bounds = i + 5
			}

			nText, err := numberInStartOfString(input[i:bounds])
			if err != nil {
				continue
			}
			if first == nil {
				n := numberMappings[*nText]
				first = &n
				last = first
			} else {
				n := numberMappings[*nText]
				last = &n
			}

		}
	}

	return sum

}

// [part1]
func firstAndLast(input string) int {
	sum := 0
	var first, last *byte
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' || i == len(input)-1 {
			if first != nil && last != nil {
				currNumber, err := strconv.Atoi(fmt.Sprint("", string(*first), string(*last)))
				fmt.Println(currNumber)
				if err != nil {
					log.Panic("first and last number are not valid")
				}
				sum += currNumber
				first, last = nil, nil

			}
			continue

		}
		if unicode.IsDigit(rune(input[i])) {
			if first == nil {
				n := input[i]
				first = &n
				last = first
			} else {
				n := input[i]
				last = &n
			}
		}
	}

	return sum

}

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(totalWinningsPart2(string(input)))
}

// Part 2
func scoreCardPart2(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		intCard, err := strconv.Atoi(string(card))
		if err != nil {
			log.Panic(err)
		}
		return intCard
	}
}

func totalWinningsPart2(input string) int {
	rows := strings.Split(input, "\n")
	players := make([]Player, 0, len(rows))
	// parses and creates a player type to store hands, bids and handTypes
	for _, row := range rows {
		handsAndBids := strings.Split(strings.TrimSpace(row), " ")
		currHand, currBid := handsAndBids[0], handsAndBids[1]
		cardCounts := make(map[rune]int)
		for _, card := range currHand {
			cardCounts[card]++
		}

		intBid, err := strconv.Atoi(currBid)
		if err != nil {
			log.Panic(err)
		}

		currPlayer := Player{
			hand:     currHand,
			handType: getHandTypePart2(cardCounts),
			bid:      intBid,
		}
		players = append(players, currPlayer)
	}

	sort.SliceStable(players, func(i int, j int) bool {
		switch {
		case players[i].handType < players[j].handType:
			return true
		case players[i].handType > players[j].handType:
			return false
		default:
			for z := 0; z < len(players[i].hand); z++ {
				if scoreCardPart2(players[i].hand[z]) != scoreCardPart2(players[j].hand[z]) {
					return scoreCardPart2(players[i].hand[z]) < scoreCardPart2(players[j].hand[z])
				}
			}
		}
		return i < j
	})

	// getting the winnings
	totalWinnings := 0
	for i, player := range players {
		fmt.Println(player.hand, player.handType)
		totalWinnings += player.bid * (i + 1)

	}
	return totalWinnings

}

func getHandTypePart2(cardCounts map[rune]int) HandType {
	handType := HandTypeHighCard
	for _, count := range cardCounts {
		switch count {
		case 5:
			return HandTypeFiveOfAKind
		case 4:
			handType = HandTypeFourOfAKind
		case 3:
			switch {
			// means we found one pair earlier
			case handType == HandTypeOnePair:
				handType = HandTypeFullHouse
			default:
				handType = HandTypeThreeOfAKind
			}
		case 2:
			switch {
			case handType == HandTypeOnePair:
				handType = HandTypeTwoPair
			case handType == HandTypeThreeOfAKind:
				handType = HandTypeFullHouse
			default:
				handType = HandTypeOnePair
			}
		}
	}

	if jokersAmount := cardCounts[rune('J')]; jokersAmount > 0 {
		switch jokersAmount {
		case 4:
			return HandTypeFiveOfAKind
		case 3:
			if handType == HandTypeFullHouse {
				return HandTypeFiveOfAKind
			}
			return HandTypeFourOfAKind
		case 2:
			if handType == HandTypeFullHouse {
				return HandTypeFiveOfAKind
			} else if handType == HandTypeTwoPair {
				return HandTypeFourOfAKind
			}
			return HandTypeThreeOfAKind
		case 1:
			switch {
			case handType == HandTypeOnePair:
				return HandTypeThreeOfAKind
			case handType == HandTypeThreeOfAKind:
				return HandTypeFourOfAKind
			case handType == HandTypeTwoPair:
				return HandTypeFullHouse
			default:
				return handType + 1
			}
		}
	}

	return handType
}

// Part 1
type HandType uint8

const (
	HandTypeHighCard HandType = iota
	HandTypeOnePair
	HandTypeTwoPair
	HandTypeThreeOfAKind
	HandTypeFullHouse
	HandTypeFourOfAKind
	HandTypeFiveOfAKind
)

type Player struct {
	hand     string
	handType HandType
	bid      int
}

func scoreCard(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		intCard, err := strconv.Atoi(string(card))
		if err != nil {
			log.Panic(err)
		}
		return intCard
	}
}

func totalWinnings(input string) int {
	rows := strings.Split(input, "\n")
	players := make([]Player, 0, len(rows))
	// parses and creates a player type to store hands, bids and handTypes
	for _, row := range rows {
		handsAndBids := strings.Split(strings.TrimSpace(row), " ")
		currHand, currBid := handsAndBids[0], handsAndBids[1]
		cardCounts := make(map[rune]int)
		for _, card := range currHand {
			cardCounts[card]++
		}

		intBid, err := strconv.Atoi(currBid)
		if err != nil {
			log.Panic(err)
		}

		currPlayer := Player{
			hand:     currHand,
			handType: getHandType(cardCounts),
			bid:      intBid,
		}
		players = append(players, currPlayer)
	}

	sort.SliceStable(players, func(i int, j int) bool {
		switch {
		case players[i].handType < players[j].handType:
			return true
		case players[i].handType > players[j].handType:
			return false
		default:
			for z := 0; z < len(players[i].hand); z++ {
				if scoreCard(players[i].hand[z]) != scoreCard(players[j].hand[z]) {
					return scoreCard(players[i].hand[z]) < scoreCard(players[j].hand[z])
				}
			}
		}
		return i < j
	})

	// getting the winnings
	totalWinnings := 0
	for i, player := range players {
		totalWinnings += player.bid * (i + 1)

	}
	return totalWinnings

}

func getHandType(cardCounts map[rune]int) HandType {
	handType := HandTypeHighCard
	for _, count := range cardCounts {
		switch count {
		case 5:
			return HandTypeFiveOfAKind
		case 4:
			return HandTypeFourOfAKind
		case 3:
			// means we found one pair earlier
			if handType == HandTypeOnePair {
				return HandTypeFullHouse
			}
			handType = HandTypeThreeOfAKind
		case 2:
			if handType == HandTypeOnePair {
				return HandTypeTwoPair
			} else if handType == HandTypeThreeOfAKind {
				return HandTypeFullHouse
			} else {
				handType = HandTypeOnePair
			}
		}
	}
	return handType
}

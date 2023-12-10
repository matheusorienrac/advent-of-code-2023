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

	fmt.Println(totalWinnings(string(input)))
}

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
	rankPlayers(players)

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
		fmt.Println(player, player.hand, player.handType)
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

func rankPlayers(players []Player) {
	// gonna sort the players by handType initially so its easier for us to check the ranks
}

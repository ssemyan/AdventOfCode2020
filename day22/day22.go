package main

import (
	"AdventOfCode2020/mods/fileload"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Read all data into string array
	lines := fileload.Fileload("day22/data.txt")

	decks := [][]int{}
	var curDeck []int
	for n, line := range lines {

		if strings.HasPrefix(line, "Player ") {
			// Make new deck
			curDeck = []int{}

		} else if line != "" {
			// Add card to current deck
			curCard, _ := strconv.Atoi(line)
			curDeck = append(curDeck, curCard)
		}

		// Save deck at end of deck or end of file
		if line == "" || n == len(lines)-1 {
			if curDeck != nil {
				decks = append(decks, curDeck)
			}
		}
	}

	fmt.Printf("Found %d players. Deck 0 size: %d \n", len(decks), len(decks[0]))

	// start playing
	var loser int
	bGameOver := false
	nRound := 0
	for {
		// When one deck is empty the game is over
		for player, deck := range decks {
			if len(deck) == 0 {
				loser = player
				bGameOver = true
				break
			}
		}
		if bGameOver {
			break
		}

		// Pop the first cards off the stacks
		card0 := decks[0][0]
		decks[0] = decks[0][1:]
		card1 := decks[1][0]
		decks[1] = decks[1][1:]

		if card0 > card1 {
			// 0 won
			decks[0] = append(decks[0], card0)
			decks[0] = append(decks[0], card1)
			fmt.Println("Player 1 won in round ", nRound)
		} else {
			// 1 won
			decks[1] = append(decks[1], card1)
			decks[1] = append(decks[1], card0)
			fmt.Println("Player 2 won in round ", nRound)
		}
		nRound++
	}

	winner := 0
	if loser == 0 {
		winner = 1
	}
	fmt.Println("Final winner: ", winner)

	// do score
	score := 0
	deckSize := len(decks[winner])
	for i := 0; i < deckSize; i++ {
		score += decks[winner][deckSize-i-1] * (i + 1)
	}
	fmt.Println("Final score: ", score)
}

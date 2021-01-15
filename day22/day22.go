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

	// start playing part one
	//partOne(decks)
	partTwo(decks)
}

func playGame(decks [][]int) int {

	nRound := 0
	prevGames := make(map[string]bool)
	for {

		// If prev seen, game goes to player 1
		if prevGame(decks, prevGames) {
			return 1 // player 2 is loser
		}

		// When one deck is empty the game is over
		for player, deck := range decks {
			if len(deck) == 0 {
				return player
			}
		}

		// Pop the first cards off the stacks
		card0 := decks[0][0]
		decks[0] = decks[0][1:]
		card1 := decks[1][0]
		decks[1] = decks[1][1:]

		// Determine if we need to recurse
		bRecuGame := false
		var innerLoser int
		if len(decks[0]) >= card0 && len(decks[1]) >= card1 {
			// Make new deck
			rdecks := [][]int{}
			deck0 := make([]int, card0)
			copy(deck0, decks[0][:card0])
			rdecks = append(rdecks, deck0)
			deck1 := make([]int, card1)
			copy(deck1, decks[1][:card1])
			rdecks = append(rdecks, deck1)
			//fmt.Printf("Playing inner round %d\n", nRound)
			innerLoser = playGame(rdecks)
			//fmt.Printf("Done with inner round %d. Player %d lost\n", nRound, innerLoser+1)
			bRecuGame = true
		}

		if (bRecuGame && innerLoser == 1) || (!bRecuGame && card0 > card1) {
			// 0 won
			decks[0] = append(decks[0], card0)
			decks[0] = append(decks[0], card1)
			//fmt.Println("Player 1 won in round ", nRound)
		} else {
			// 1 won
			decks[1] = append(decks[1], card1)
			decks[1] = append(decks[1], card0)
			//fmt.Println("Player 2 won in round ", nRound)
		}
		nRound++
	}
}

func partTwo(decks [][]int) {

	loser := playGame(decks)
	winner := 0
	if loser == 0 {
		winner = 1
	}
	fmt.Println("Part One Final winner: ", winner)

	// do score
	score := 0
	deckSize := len(decks[winner])
	for i := 0; i < deckSize; i++ {
		score += decks[winner][deckSize-i-1] * (i + 1)
	}
	fmt.Println("Part One Final score: ", score)
}

func prevGame(decks [][]int, prevGames map[string]bool) bool {

	// Make deck into string
	decksStr := ""
	for _, deck := range decks {
		for _, card := range deck {
			decksStr += fmt.Sprintf("%d", card) + "."
		}
		decksStr += "|" // be sure to mark between the two decks
	}
	_, exists := prevGames[decksStr]
	if exists {
		return true
	}
	prevGames[decksStr] = true
	return false
}

func partOne(decks [][]int) {
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
	fmt.Println("Part One Final winner: ", winner)

	// do score
	score := 0
	deckSize := len(decks[winner])
	for i := 0; i < deckSize; i++ {
		score += decks[winner][deckSize-i-1] * (i + 1)
	}
	fmt.Println("Part One Final score: ", score)
}

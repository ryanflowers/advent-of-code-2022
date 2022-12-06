package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Guide

// Column 1(Opponent)
// A for Rock, B for Paper, and C for Scissors.

// Column 2(You)
// X for Rock, Y for Paper, and Z for Scissors

// Total
// Your total score is the sum of your scores for each round.
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)

var opponentShapeMap = map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
var mineShapeMap = map[string]string{"X": "Rock", "Y": "Paper", "Z": "Scissors"}

func getInput() (string, error) {

	content, err := ioutil.ReadFile("./day 2/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the content to type string
	sb := string(content)

	return sb, nil
}

func toShapeValue(shape string) int {
	if shape == "Rock" {
		return 1
	} else if shape == "Paper" {
		return 2
	} else {
		// Scissors
		return 3
	}
}

func main() {

	var guide, err = getInput()
	if err != nil {
		panic(err)
	}

	moves := strings.Split(guide, "\n")

	log.Printf(moves[0])

	total := 0
	for _, move := range moves {
		// fmt.Println(i, move)

		// New move
		splitMove := strings.Split(move, " ")

		opponentShape := opponentShapeMap[splitMove[0]]
		fmt.Printf("Opponent Shape %v", opponentShape)
		fmt.Println()

		mineShape := mineShapeMap[splitMove[1]]
		fmt.Printf("Mine Shape %v", mineShape)
		fmt.Println()

		opponentValue := toShapeValue(opponentShape)
		fmt.Printf("Opponent Value %v", opponentValue)
		fmt.Println()

		mineValue := toShapeValue(mineShape)
		fmt.Printf("Mine Value %v", mineValue)
		fmt.Println()

		total += mineValue

		// Determine winner
		if opponentValue == mineValue {
			// Draw
			fmt.Printf("Draw")

			total += 3
		} else if (mineValue > opponentValue ||
			mineShape == "Rock" && opponentShape == "Scissors") &&
			!(opponentShape == "Rock" && mineShape == "Scissors") {

			fmt.Printf("Win")

			// Win
			total += 6
		} else {
			//Lost so nothing
			fmt.Printf("Lost")
		}

		fmt.Println()

	}

	fmt.Printf("Total %d", total)
	fmt.Println()

}

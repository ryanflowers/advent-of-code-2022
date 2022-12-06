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
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.

// Total
// Your total score is the sum of your scores for each round.
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)

var opponentShapeMap = map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
var moveResultMap = map[string]string{"X": "Lose", "Y": "Draw", "Z": "Win"}

func getInput() (string, error) {

	content, err := ioutil.ReadFile("./day 2/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the content to type string
	sb := string(content)

	return sb, nil
}

func toResultValue(result string) int {
	if result == "Win" {
		return 6
	} else if result == "Draw" {
		return 3
	} else {
		// Lost
		return 0
	}
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

		moveResult := moveResultMap[splitMove[1]]
		fmt.Printf("Move Result %v", moveResult)
		fmt.Println()

		resultValue := toResultValue(moveResult)
		fmt.Printf("Move Result Value %v", resultValue)
		fmt.Println()

		total += resultValue

		// Determine mine move value
		if moveResult == "Draw" {
			// Draw
			fmt.Printf("Draw")

			total += toShapeValue(opponentShape)
		} else if moveResult == "Win" {
			// Win
			fmt.Printf("Win")

			if opponentShape == "Rock" {
				total += toShapeValue("Paper")
			} else if opponentShape == "Paper" {
				total += toShapeValue("Scissors")
			} else {
				total += toShapeValue("Rock")
			}
		} else {
			//Lost
			fmt.Printf("Lost")

			if opponentShape == "Rock" {
				total += toShapeValue("Scissors")
			} else if opponentShape == "Paper" {
				total += toShapeValue("Rock")
			} else {
				total += toShapeValue("Paper")
			}
		}

		fmt.Println()

	}

	fmt.Printf("Total %d", total)
	fmt.Println()

}

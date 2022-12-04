package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func getInput() (string, error) {

	content, err := ioutil.ReadFile("day-1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the content to type string
	sb := string(content)

	return sb, nil
}

func main() {
	fmt.Println("Hello, World!")

	var calories, err = getInput()
	if err != nil {
		panic(err)
	}

	splitCalories := strings.Split(calories, "\n")
	// Traverse each elf cals and sum keep track of the largest

	log.Printf(splitCalories[0])

	length := len(splitCalories)
	maxCalories := make([]int, length)
	currentCalories := 0
	numberElves := 0
	for i, s := range splitCalories {
		fmt.Println(i, s)

		// New elf
		if s != "" {
			// Add
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			currentCalories += i

		} else {

			maxCalories[numberElves] = currentCalories

			currentCalories = 0

			numberElves++
		}
	}

	// Could do this without keeping all and sorting but works for now

	sort.Ints(maxCalories)

	// Sum the top 3

	total := maxCalories[length-1] + maxCalories[length-2] + maxCalories[length-3]

	fmt.Printf("Max calories %d", total)
	fmt.Println()
	fmt.Printf("Number of elves %d", numberElves)
}

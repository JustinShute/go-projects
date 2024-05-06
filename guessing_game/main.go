package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lowNumber int = 1
const highNumber int = 100

var attempts uint = 1

func main() {

	systemPick := pickRandomNumber()

	for {
		userGuess, isValid := userInput()
		if isValid != nil {
			continue
		} else if userGuess < lowNumber || userGuess > highNumber {
			continue
		} else if userGuess == systemPick {
			if attempts == 1 {
				fmt.Printf("Correct, the number is %v, you guessed the number in %v attempt!\n", systemPick, attempts)
			} else {
				fmt.Printf("Correct, the number is %v, you guessed the number in %v attempts!\n", systemPick, attempts)
			}
			break
		} else if userGuess > systemPick {
			fmt.Println("Wrong, your guess is too high, try a lower number.")
		} else {
			fmt.Println("Wrong, your guess is too low, try a higher number.")
		}
		attempts++
	}
}

func userInput() (int, error) {
	var userGuess int
	fmt.Printf("Please pick a number between %v and %v.\n", lowNumber, highNumber)
	if _, err := fmt.Scan(&userGuess); err != nil {
		fmt.Println("Invalid entry, please enter a valid number.")
		return userGuess, err
	}
	return userGuess, nil
}

func pickRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(highNumber-lowNumber) + lowNumber
}

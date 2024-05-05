package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	userWins := 0
	systemWins := 0
	fmt.Println("Welcome to my rock, paper, scissors application.")
	for {
		fmt.Printf("The score is you %v and me %v.\n", userWins, systemWins)
		userPick := userChoice()
		systemPick := systemChoice()

		if userPick == systemPick {
			fmt.Printf("We both picked %v, lets play again!\n", systemPick)
			continue
		} else if userPick == "rock" && systemPick == "paper" ||
			userPick == "paper" && systemPick == "scissors" ||
			userPick == "scissors" && systemPick == "rock" {
			fmt.Printf("I win, you picked %v and I picked %v\n", userPick, systemPick)
			systemWins++
			continue
		} else if userPick == "rock" && systemPick == "scissors" ||
			userPick == "paper" && systemPick == "rock" ||
			userPick == "scissors" && systemPick == "paper" {
			fmt.Printf("You win, you picked %v and I picked %v.\n", userPick, systemPick)
			userWins++
			continue

		}
	}
}

func userChoice() string {
	var userSelection string
	fmt.Println("Please pick rock, paper, or scissors.")
	for {
		fmt.Scanln(&userSelection)
		userSelection = strings.ToLower(userSelection)
		if userSelection == "rock" || userSelection == "paper" || userSelection == "scissors" {
			fmt.Printf("You selected %v.\n", userSelection)
			return userSelection
		} else {
			fmt.Println("Invalid entry, please select either rock, paper or scissors.")
			continue
		}

	}

}

func systemChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(choices))
	randomChoice := choices[randomIndex]
	return randomChoice
}

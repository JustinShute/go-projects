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

		userPick := userChoice()
		systemPick := systemChoice()

		if userPick == systemPick {
			fmt.Printf("We both picked %v, lets play again!\n", systemPick)
			fmt.Printf("The score is you %v and me %v.\n", userWins, systemWins)
			isPlayAgain := playAgain()
			if isPlayAgain {
				continue
			} else {
				return
			}
		} else if userPick == "rock" && systemPick == "paper" ||
			userPick == "paper" && systemPick == "scissors" ||
			userPick == "scissors" && systemPick == "rock" {
			fmt.Printf("I win, you picked %v and I picked %v\n", userPick, systemPick)
			systemWins++
			fmt.Printf("The score is you %v and me %v.\n", userWins, systemWins)
			isPlayAgain := playAgain()
			if isPlayAgain {
				continue
			} else {
				return
			}
		} else {
			fmt.Printf("You win, you picked %v and I picked %v.\n", userPick, systemPick)
			userWins++
			fmt.Printf("The score is you %v and me %v.\n", userWins, systemWins)
			isPlayAgain := playAgain()

			if isPlayAgain {
				continue
			} else {
				return
			}

		}
	}
}

func userChoice() string {
	var userSelection string
	fmt.Println("Please pick rock, paper, or scissors: ")
	for {
		fmt.Scanln(&userSelection)
		userSelection = strings.ToLower(userSelection)
		if userSelection == "rock" || userSelection == "paper" || userSelection == "scissors" {

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

func playAgain() bool {
	var isPlayAgain bool
	var userPlayAgain string
	fmt.Println("Do you want to play again?")
	fmt.Scanln(&userPlayAgain)
	switch userPlayAgain {
	case "y", "yes":
		isPlayAgain = true
		return isPlayAgain
	case "n", "no":
		isPlayAgain = false
		return isPlayAgain
	default:
		fmt.Println("Please pick a valid response")
		return playAgain()
	}
}

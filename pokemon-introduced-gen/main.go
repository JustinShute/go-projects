package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// the APIs for each generation of Pokemon
var apiUrls = []string{
	"https://pokeapi.co/api/v2/generation/1/",
	"https://pokeapi.co/api/v2/generation/2/",
	"https://pokeapi.co/api/v2/generation/3/",
	"https://pokeapi.co/api/v2/generation/4/",
	"https://pokeapi.co/api/v2/generation/5/",
	"https://pokeapi.co/api/v2/generation/6/",
	"https://pokeapi.co/api/v2/generation/7/",
	"https://pokeapi.co/api/v2/generation/8/",
	"https://pokeapi.co/api/v2/generation/9/",
}

// struct used to pull information from the APIs
type GenerationResponse struct {
	PokemonSpecies []struct {
		PokemonName string `json:"name"`
	} `json:"pokemon_species"`
}

// maxes the interations of the APIs to the amount of APIs in the slice, if the Pokemon isnt found in any API, print a statement.
func fetchPokemonCheckUserInput(urls []string, userInput string, index int) {
	if index >= len(urls) {
		fmt.Println("Please input a valid Pokemon.")
		return
	}

	//handle error for if the Pokemon isnt found in that iteration, then increment the API index up 1, run the loop again.
	resp, err := http.Get(urls[index])
	if err != nil {
		fmt.Println("Error: ", err)
		fetchPokemonCheckUserInput(urls, userInput, index+1)
	}
	defer resp.Body.Close()

	//decode the information provided by the API.
	var apiResponse GenerationResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Error: ", err)
		fetchPokemonCheckUserInput(urls, userInput, index+1)
		return
	}

	//check if user input is in an API, if it is print a statement with the index +1, since the generation is 1 off of the index.
	for _, species := range apiResponse.PokemonSpecies {
		if species.PokemonName == strings.ToLower(userInput) {
			fmt.Printf("%v was introduced in Generation %v.\n", strings.Title(species.PokemonName), index+1)
			return

		}
	}

	fetchPokemonCheckUserInput(urls, userInput, index+1)
}

// user picks a pokemon, then program is called starting on API #1.
func main() {
	var userPick string
	fmt.Println("Please pick a Pokemon: ")
	fmt.Scanln(&userPick)

	fetchPokemonCheckUserInput(apiUrls, userPick, 0)
}

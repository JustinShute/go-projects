package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var userResponse string
var url string = "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=200000000000000"

type apiFormat struct {
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func UserInput() string {
	fmt.Println("Please enter the Pokemon's name: ")
	fmt.Scan(&userResponse)
	return userResponse
}

func ApiResponse() ([]string, []string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())

	}
	defer resp.Body.Close()

	var apiInformation apiFormat
	err = json.NewDecoder(resp.Body).Decode(&apiInformation)
	if err != nil {
		fmt.Println("Error: ", err.Error())

	}

	var PokemonName []string
	var PokemonUrl []string
	for _, pkmnName := range apiInformation.Results {
		PokemonName = append(PokemonName, pkmnName.Name)

	}

	for _, pkmnUrl := range apiInformation.Results {
		PokemonUrl = append(PokemonUrl, pkmnUrl.Url)
	}

	return PokemonName, PokemonUrl
}

func apiForUserInput(resUser string, pokemonName []string, pokemonUrl []string, index int) string {
	if index >= len(pokemonName) {
		fmt.Println("Please enter a valid Pokemon name.")
		return ""
	}
	for ind, name := range pokemonName {
		if name == resUser {
			return pokemonUrl[ind]
		}

	}
	apiForUserInput(resUser, pokemonName, pokemonUrl, index+1)
	return ""
}

//func variables() string {
//ResultOfUser := strings.ToLower(UserInput())

//pokemonName, pokemonUrl := ApiResponse()
//PokemonUrl := apiForUserInput(ResultOfUser, pokemonName, pokemonUrl, 0)
//if PokemonUrl != "" {
//fmt.Printf("API for %v is: %v\n", strings.Title(ResultOfUser), pokemonUrl)
//} else {
//variables()
//}
//return ResultOfUser
//}

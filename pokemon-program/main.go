package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type apiFormat struct {
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type Stat struct {
	Name  string `json:"name"`
	Value int    `json:"base_stat"`
}

type PokemonStats struct {
	Stats []Stat `json:"stats"`
}

type PokemonTypes struct {
	Types []struct {
		PkmnType struct {
			TypeName string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func userInput() string {
	var userResponse string
	fmt.Print("Please enter the Pokemon's name: ")
	fmt.Scan(&userResponse)
	return userResponse
}

func apiResponse() ([]string, []string) {
	var url string = "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=200000000000000"
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

	var pokemonNames []string
	var pokemonUrl []string
	for _, pkmnName := range apiInformation.Results {
		pokemonNames = append(pokemonNames, pkmnName.Name)

	}

	for _, pkmnUrl := range apiInformation.Results {
		pokemonUrl = append(pokemonUrl, pkmnUrl.Url)
	}

	return pokemonNames, pokemonUrl
}

func apiForUserInput(resUser string, pokemonNames []string, pokemonUrl []string, index int) string {
	if index >= len(pokemonNames) {
		fmt.Println("Invalid entry...")
		main()
	}
	for i, name := range pokemonNames {
		if name == strings.ToLower(resUser) {
			return pokemonUrl[i]
		}

	}
	apiForUserInput(resUser, pokemonNames, pokemonUrl, index+1)
	return ""
}

func requestType(url string) []string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	defer res.Body.Close()

	var apiInformation PokemonTypes
	err = json.NewDecoder(res.Body).Decode(&apiInformation)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	var pkmnType []string
	for _, data := range apiInformation.Types {
		pkmnType = append(pkmnType, strings.Title(data.PkmnType.TypeName))
	}
	return pkmnType
}

func requestStats(pokemonUrl string) []int {

	res, err := http.Get(pokemonUrl)
	if err != nil {
		fmt.Print("Error: ", err.Error())
	}
	defer res.Body.Close()

	var apiInformation PokemonStats
	err = json.NewDecoder(res.Body).Decode(&apiInformation)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	var statsOfPokemon []int

	for _, data := range apiInformation.Stats {
		statsOfPokemon = append(statsOfPokemon, data.Value)
	}
	return statsOfPokemon
}

func printPokemonData(url string, responseFromUser string) {
	pkmnStats := requestStats(url)
	var pkmnStatsTotal int = pkmnStats[0] + pkmnStats[1] + pkmnStats[2] + pkmnStats[3] + pkmnStats[4] + pkmnStats[5]
	pkmnType := requestType(url)

	if len(pkmnType) == 1 {
		fmt.Printf("\nName: %v\n\n   HP: %v\n  Att: %v\n  Def: %v\nSpAtt: %v\nSpDef: %v\nSpeed: %v\nTotal: %v\n\nType: %v\n\n", strings.Title(responseFromUser), pkmnStats[0], pkmnStats[1], pkmnStats[2], pkmnStats[3], pkmnStats[4], pkmnStats[5], pkmnStatsTotal, pkmnType[0])
	} else {
		fmt.Printf("\nName: %v\n\n   HP: %v\n  Att: %v\n  Def: %v\nSpAtt: %v\nSpDef: %v\nSpeed: %v\nTotal: %v\n\nType: %v / %v\n\n", strings.Title(responseFromUser), pkmnStats[0], pkmnStats[1], pkmnStats[2], pkmnStats[3], pkmnStats[4], pkmnStats[5], pkmnStatsTotal, pkmnType[0], pkmnType[1])
	}
}

func main() {
	responseFromUser := strings.ToLower(userInput())
	nameSlice, urlSlice := apiResponse()
	pokemonUrl := apiForUserInput(responseFromUser, nameSlice, urlSlice, 0)

	printPokemonData(pokemonUrl, responseFromUser)
}

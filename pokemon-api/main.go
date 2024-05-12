package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var url string = "https://pokeapi.co/api/v2/pokemon/21/"

type Stat struct {
	Name  string `json:"name"`
	Value int    `json:"base_stat"`
}

type PokemonStats struct {
	Stats []Stat `json:"stats"`
}

func apiRequest() []int {

	res, err := http.Get(url)
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

func main() {
	pkmnStats := apiRequest()
	if len(pkmnStats) == 0 {
		return
	}
	pkmnStatsTotal := pkmnStats[0] + pkmnStats[1] + pkmnStats[2] + pkmnStats[3] + pkmnStats[4] + pkmnStats[5]
	fmt.Printf("   HP: %v\n  Att: %v\n  Def: %v\nSpAtt: %v\nSpDef: %v\nSpeed: %v\nTotal: %v\n", pkmnStats[0], pkmnStats[1], pkmnStats[2], pkmnStats[3], pkmnStats[4], pkmnStats[5], pkmnStatsTotal)
}

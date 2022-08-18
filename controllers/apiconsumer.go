package controllers

import (
	"encoding/json"
	"fmt"
	"golang-crud-sql/database"
	"golang-crud-sql/entities"
	"net/http"
)

func ApiConsumer() {
	//recebe os dados da API no formato de []bytes
	response, _ := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	var responseObject entities.Response

	json.NewDecoder(response.Body).Decode(&responseObject)

	var pokemonNames entities.NamesDataBase
	var pokemonsTypes entities.TypesDataBase

	for index, value := range responseObject.Pokemon {
		pokemonNames.ID = value.EntryNo
		pokemonNames.Name = value.Species.Name

		database.Instance.Create(&pokemonNames)

		pokemonsTypes.NamesDataBaseID = value.EntryNo

		response, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", index+1))
		var responseObject entities.PokemonDetail

		json.NewDecoder(response.Body).Decode(&responseObject)

		for _, value := range responseObject.PokemonTypesData {
			pokemonsTypes.PokemonTypes = value.PokemonTypes.Name

			database.Instance.Create(&pokemonsTypes)
		}

	}

}

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

	var pokemon entities.PokemonDataBase

	for index, value := range responseObject.Pokemon {
		pokemon.ID = value.EntryNo
		pokemon.Name = value.Species.Name

		response, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", index+1))

		var responseObject entities.PokemonDetail

		json.NewDecoder(response.Body).Decode(&responseObject)

		//acredito que ta errado aqui nessa posição nunca zera e só cresce
		//confirar se colocando no final do for o nil zera a variavel
		var typesData string

		for _, value := range responseObject.PokemonTypesData {
			//tentar o slot para criar no struct
			typesData = typesData + value.Types.Name

		}
		pokemon.Type = typesData
		database.Instance.Create(&pokemon)
		typesData = ""
	}
	// for i := 0; i < len(responseObject.Pokemon); i++ {
	// 	fmt.Println(responseObject.Pokemon[i].Species.Name)
	// }
}

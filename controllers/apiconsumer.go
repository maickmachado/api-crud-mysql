package controllers

import (
	"encoding/json"
	"fmt"
	"golang-crud-sql/database"
	"golang-crud-sql/entities"
	"net/http"
	"os"
)

func ApiConsumer() {
	//recebe os dados da API no formato de []bytes
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var responseObject entities.Response

	json.NewDecoder(response.Body).Decode(&responseObject)

	var pokemon entities.PokemonDataBase

	for _, value := range responseObject.Pokemon {
		pokemon.ID = value.EntryNo
		pokemon.Name = value.Species.Name
		database.Instance.Create(&pokemon)
	}
	// for i := 0; i < len(responseObject.Pokemon); i++ {
	// 	fmt.Println(responseObject.Pokemon[i].Species.Name)
	// }
}

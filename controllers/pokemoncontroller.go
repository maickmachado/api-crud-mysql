package controllers

import (
	"encoding/json"
	"golang-crud-sql/database"
	"golang-crud-sql/entities"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func CreateDataPokemon(w http.ResponseWriter, r *http.Request) {
	//define o estilo que os dados serão mostrados em w
	w.Header().Set("Content-Type", "application/json")
	//cria uma variável do tipo Product do pacote entities
	var pokemon entities.NamesDataBase
	//decodifica os dados recebidos em 'r' e coloca dentro do endereço na memória da variável product
	//pega dados no estilo json para struct
	json.NewDecoder(r.Body).Decode(&pokemon)
	//usando GORM é criado um novo produto usando os dados da variável product
	database.Instance.Create(&pokemon)
	//codifica os dados do produto no estilo struct para json
	//verificar se só serve para printar a informação em w
	json.NewEncoder(w).Encode(pokemon)
}

func checkIfPokemonExists(pokemonId string) bool {
	var pokemon entities.NamesDataBase
	//passar como parâmetro o endereço de memória da variável e segundo parametro o ID
	//usa GORM para achar o primeiro item
	database.Instance.First(&pokemon, pokemonId)
	//se não tiver dados o GORM retorna o valor 0
	return pokemon.ID != 0
	// if product.ID == 0 {
	// 	return false
	// }
	// return true
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	//utiliza o ID passado na URL host/api/products/{id}
	//usa o MUX para extrair o ID da URL recebida no 'r' e associar a variável criada productId
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		//faz onde, o que
		//no w escreve a frase em formato json (?)
		json.NewEncoder(w).Encode("Pokemon não encontrado!")
		return
	}
	var pokemon entities.NamesDataBase
	//utiliza o endereço na memória do product para procurar o primeiro item com o productId
	database.Instance.First(&pokemon, pokemonId)
	w.Header().Set("Content-Type", "application/json")
	//codifica product e envia para w
	json.NewEncoder(w).Encode(pokemon)
}

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	pokemonName := mux.Vars(r)["name"]
	if !checkIfPokemonExists(pokemonName) {
		//faz onde, o que
		//no w escreve a frase em formato json (?)
		json.NewEncoder(w).Encode("Pokemon não encontrado!")
		return
	}
}

func GetDataPokemons(w http.ResponseWriter, r *http.Request) {
	//criado um slice do struct Product
	var pokemon []entities.NamesDataBase
	//procura no banco de dados todos os products
	//database.Instance.Find(&pokemon)
	database.Instance.Preload(clause.Associations).Find(&pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
}

func UpdateDataPokemon(w http.ResponseWriter, r *http.Request) {
	//utiliza o ID passado na URL host/api/products/{id}
	//usa o MUX para extrair o ID da URL recebida no 'r' e associar a variável criada productId
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		//faz onde, o que
		//no w escreve a frase em formato json (?)
		json.NewEncoder(w).Encode("Pokemon não encontrado!")
		return
	}
	var pokemon entities.NamesDataBase
	//pega o primeiro item no banco de dados com o determinado ID
	database.Instance.First(&pokemon, pokemonId)
	//decodifica o dado recebido em 'r' no tipo product
	json.NewDecoder(r.Body).Decode(&pokemon)
	//usa o GORM para salvar no banco de dados o tipo decodificado
	database.Instance.Save(&pokemon)
	w.Header().Set("Content-Type", "application/json")
	//codifico o product e envio para 'w'
	json.NewEncoder(w).Encode(pokemon)
}

func DeleteDataPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//utiliza o ID passado na URL host/api/products/{id}
	//usa o MUX para extrair o ID da URL recebida no 'r' e associar a variável criada productId
	pokemonId := mux.Vars(r)["id"]
	if !checkIfPokemonExists(pokemonId) {
		w.WriteHeader(http.StatusNotFound)
		//faz onde, o que
		//no w escreve a frase em formato json (?)
		json.NewEncoder(w).Encode("Pokemon não encontrado!")
		return
	}
	var pokemon entities.NamesDataBase
	//GORM acessa o banco de dados e deleta o product
	database.Instance.Delete(&pokemon, pokemonId)
	json.NewEncoder(w).Encode("Pokemon deletado com sucesso!")
}

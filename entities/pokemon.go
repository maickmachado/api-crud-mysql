package entities

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

//Structs dos tipos dos pokemons
type PokemonDetail struct {
	PokemonTypesData []PokemonTypesData `json:"types"`
}

type PokemonTypesData struct {
	PokemonTypes PokemonTypes `json:"type"`
}

type PokemonTypes struct {
	Name string `json:"name"`
}

//struct do database com os nomes
type NamesDataBase struct {
	//gorm.Model
	ID            int             `json:"id"`
	Name          string          `json:"pokemon"`
	TypesDataBase []TypesDataBase `json:"types"`
}

//struct do database com os tipos
type TypesDataBase struct {
	//gorm.Model
	NamesDataBaseID int    `json:"-"`
	PokemonTypes    string `json:"type"`
	//NamesDataBase   NamesDataBase
}

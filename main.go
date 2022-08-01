package main

import (
	"fmt"
	"golang-crud-sql/config"
	"golang-crud-sql/database"
	"golang-crud-sql/routes"
	"log"
)

func main() {
	//inicializa a função LoadAppConfig do arquivo config.go
	//utiliza o pacote viper para otimizar as configurações
	config.LoadAppConfig()

	//acessa o arquivo routes.go e inicializa a função HandleRequest
	routes.HandleRequest()

	//acessa a variavel ApConfig do arquivo config.go do pacote main e acessa o conection string
	//sing the database package and the connection string from the JSON file, the application attempts to connect to the MySQL Database
	//and migrates the required table.
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
}

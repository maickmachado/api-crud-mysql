package main

import (
	"fmt"
	"golang-crud-sql/config"
	"golang-crud-sql/database"
	"golang-crud-sql/routes"
)

func main() {
	//inicializa a função LoadAppConfig do arquivo config.go
	//utiliza o pacote viper para otimizar as configurações
	config.LoadAppConfig()
	fmt.Println("1..")
	//acessa a variavel ApConfig do arquivo config.go do pacote main e acessa o conection string
	//sing the database package and the connection string from the JSON file, the application attempts to connect to the MySQL Database
	//and migrates the required table.
	fmt.Println(config.AppConfig.ConnectionString)
	database.Connect(config.AppConfig.ConnectionString)
	fmt.Println("chegou aqui")
	database.Migrate()
	//acessa o arquivo routes.go e inicializa a função HandleRequest
	routes.HandleRequest()
	fmt.Println("2..")

	// Start the server
	//log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
}

package main

import (
	"golang-crud-sql/config"
	"golang-crud-sql/database"
	"golang-crud-sql/routes"
)

func main() {
	//inicializa a função LoadAppConfig do pacote config
	//utiliza o pacote viper para otimizar as configurações
	config.LoadAppConfig()
	//acessa a função Connect no pacote database
	//a função recebe o connection string do pacote config acessado através da variável AppConfig
	//nomedousuári:senha@protocolo(127.0.0.1:3306)/nomedobancodedados
	//"connection_string": "root:abc@tcp(127.0.0.1:3306)/apimysql",
	database.Connect(config.AppConfig.ConnectionString)
	//acessa a função Migrate do pacote database
	database.Migrate()
	//log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	//acessa a função HandleRequest do pacote routes
	routes.HandleRequest()
}

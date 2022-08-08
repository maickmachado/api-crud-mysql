package config

import (
	"log"

	"github.com/spf13/viper"
)

//struct que guarda todas as configurações
type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

//variável que do tipo Config que poderá ser acessada de outros arquivos e pacotes
var AppConfig *Config

//Viper é usado para carregar as configurações do arquivo config.json
//os valores são passados para a variável criada AppConfig
//a função LoadAppConfig será chamada da função main
func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	//caminho de onde está o arquivo
	viper.AddConfigPath("./config")
	//nome do arquivo que queremos carregar
	viper.SetConfigName("config")
	//extensão do arquivo
	viper.SetConfigType("json")
	//lê o arquivo e carrega seu conteúdo
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	//decodifica o json para o tipo struct AppConfig
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

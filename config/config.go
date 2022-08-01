package config

import (
	"log"

	"github.com/spf13/viper"
)

//Here we define a struct that will contain the allowed configurations.
//In our case, it will be the port number and the MySQL connection string.
type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

//Here we define the AppConfig variable that will be accessed by other files and packages within the application code.
var AppConfig *Config

//Here is where we use Viper to load configurations from the config.json file ( which we will create in the next step)
//and assign its values to the AppConfig variable.
//The idea will be to call the LoadAppConfig function from the main program which in turn will be loading the data from
//the JSON file into the AppConfig variable.
func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	// caminho alternativo onde está o arquivo
	viper.AddConfigPath("./config")
	// nome do arquivo que queremos carregar
	viper.SetConfigName("config")
	// extensão do arquivo
	viper.SetConfigType("json")
	// lê o arquivo e carrega seu conteúdo
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

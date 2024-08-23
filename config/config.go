package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	TestConnectionString string `mapstructure:"test_connection_string"`
}

var AppConfig *Config

func LoadAppConfig() {
	
	log.Println("⏳ Carregando as Variáveis de Ambiente...")

	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("❌ Erro de Leitura! \n", err.Error())
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("❌ Falha ao Carregar! \n", err.Error())
	}
}

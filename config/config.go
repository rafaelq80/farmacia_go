package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Secret string `mapstructure:"secret"`
	SenderEmail string `mapstructure:"sender_email"`
	SmtpHost string `mapstructure:"smtp_host"`
	SmtpPort int `mapstructure:"smtp_port"`
	SmtpUser string `mapstructure:"smtp_user"`
	SmtpPassword string `mapstructure:"smtp_password"`
}

var AppConfig *Config

func LoadAppConfig(arquivo string) {
	
	log.Println("⏳ Carregando as Variáveis de Ambiente...")

	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName(arquivo)

	log.Printf("Arquivo: %s", arquivo)
	
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("❌ Erro de Leitura! \n", err.Error())
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("❌ Falha ao Carregar! \n", err.Error())
	}
}

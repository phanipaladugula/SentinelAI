package config

import(
	"log"
	"github.com/spf13/viper"
)

type Config struct{
	Port string
}

func LoadConfig() *Config{

	viper.SetConfigFile(".env")
	err:=viper.ReadInConfig()
	
	if err!=nil{
		log.Fatal("Cannot read env file")
	}

	cfg := &Config{
		Port: viper.GetString("PORT"),
	}
	return cfg
}
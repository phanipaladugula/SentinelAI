package config

import(
	"log"
	"github.com/spf13/viper"
)

type Config struct{
	Port string

	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func LoadConfig() *Config{

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err:=viper.ReadInConfig()
	
	if err!=nil{
		log.Fatal("Cannot read env file")
	}

	cfg := &Config{
		Port: viper.GetString("PORT"),
		DBHost : viper.GetString("DB_HOST"),
		DBPort : viper.GetString("DB_PORT"),
		DBUser : viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName: viper.GetString("DB_NAME"),
	}

	return cfg
}
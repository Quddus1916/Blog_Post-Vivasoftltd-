package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port      string
	SecretKey string
}

func Initconfig() Config {
	godotenv.Load()
	var config Config
	config.Port = os.Getenv("PORT")
	config.SecretKey = os.Getenv("SECRETKEY")

	return config

}

func Getconfig() Config {
	return Initconfig()
}

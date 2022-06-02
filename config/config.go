package config

import (
	//"github.com/joho/godotenv"
	"github.com/spf13/viper"
	//"os"
	"fmt"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRETKEY"`
}

func Initconfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("App")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = viper.Unmarshal(&config)
	return
}

func Getconfig() Config {
	config, _ := Initconfig()
	return config

}

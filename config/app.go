package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
	Database          Database
	OpenExchangeRates OpenExchangeRates
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type OpenExchangeRates struct {
	AppId string
	Url   string
}

func LoadConfig() *App {
	var appConfig App
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetEnvPrefix("ki")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(fmt.Errorf("cannot unmarshal config: %v", err))
	}
	return &appConfig
}

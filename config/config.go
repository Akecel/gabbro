package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitConfig set all configuration for the app
func InitConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}

func InitTestConfig() {
	viper.SetConfigFile("../.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}
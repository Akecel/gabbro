package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/bin/")
	viper.SetConfigName("gabbro")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}

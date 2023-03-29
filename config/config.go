package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath(".") // Local config
	viper.AddConfigPath("$HOME/bin/") // Binary config
	viper.AddConfigPath("../") // Test config
	viper.SetConfigName("gabbro")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}

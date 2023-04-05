package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Local config path
	viper.AddConfigPath(".")

	// Binary config path
	viper.AddConfigPath("$GOPATH/bin/")
	viper.AddConfigPath("$HOME/bin/")
	viper.AddConfigPath("%USERPROFILE%/bin/")

	// Test config path
	viper.AddConfigPath("../")
	
	viper.SetConfigName("gabbro")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}

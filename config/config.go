package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

func InitConfig() {
	var binPath string
	if runtime.GOOS == "windows" {
		binPath = os.Getenv("USERPROFILE") + "/bin"
	} else {
		binPath = os.Getenv("HOME") + "/bin"
	}

	// Local config path
	viper.AddConfigPath(".")

	// Binary config path
	viper.AddConfigPath("$GOPATH/bin")
	viper.AddConfigPath(binPath)

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

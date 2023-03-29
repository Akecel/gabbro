package config

import (
	"flag"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/spf13/viper"
)

func InitClient() *igdb.Client {
	InitConfig()

	clientId := viper.GetString("client-id")
	accessToken := viper.GetString("access-token")

	client := igdb.NewClient(clientId, accessToken, nil)

	return client
}

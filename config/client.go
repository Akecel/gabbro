package config

import (
	"github.com/spf13/viper"
	"github.com/Henry-Sarabia/igdb/v2"
)

func InitClient() (*igdb.Client) {
	clientId := viper.GetString("client-id")
	accessToken := viper.GetString("access-token")

	client := igdb.NewClient(clientId, accessToken, nil)

	return client
}
package config

import (
	"github.com/spf13/viper"
	"github.com/Henry-Sarabia/igdb/v2"
)

func InitClient() (*igdb.Client) {
	clientId := viper.GetString("CLIENT_ID")
	accessToken := viper.GetString("ACCESS_TOKEN")

	client := igdb.NewClient(clientId, accessToken, nil)

	return client
}
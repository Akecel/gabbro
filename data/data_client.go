package data

import (
	"log"
	
	"github.com/akecel/gabbro/config"

	"github.com/Henry-Sarabia/igdb/v2"
)

var client *igdb.Client

func init() {
	client = config.InitClient()
}

func GetGameData(name string, limit int) ([]*igdb.Game) {
	game, err := client.Games.Search(
		name,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return game
}
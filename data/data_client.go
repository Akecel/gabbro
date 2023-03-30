package data

import (
	"log"

	"github.com/akecel/gabbro/config"

	"github.com/Henry-Sarabia/igdb/v2"
)

var (
	client = config.InitClient()
)

func GetGamesDataByName(name string, limit int) []*igdb.Game {
	games, err := client.Games.Search(
		name,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return games
}

func GetGamesDataByIDs(ids []int, limit int) []*igdb.Game {
	games, err := client.Games.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return games
}

func GetGenresDataByIDs(ids []int, limit int) []*igdb.Genre {
	genres, err := client.Genres.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return genres
}

func GetThemesDataByIDs(ids []int, limit int) []*igdb.Theme {
	themes, err := client.Themes.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return themes
}

func GetPlatformsDataByIDs(ids []int, limit int) []*igdb.Platform {
	platforms, err := client.Platforms.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return platforms
}

func GetInvolvedCompaniesDataByIDs(ids []int, limit int) []*igdb.InvolvedCompany {
	involvedCompanies, err := client.InvolvedCompanies.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return involvedCompanies
}

func GetCompaniesDataByID(id int, limit int) *igdb.Company {
	involvedCompanies, err := client.Companies.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return involvedCompanies
}

func GetCoverDataByIDs(id int, limit int) *igdb.Cover {
	cover, err := client.Covers.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return cover
}

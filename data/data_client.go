package data

import (
	"log"

	"github.com/akecel/gabbro/config"

	"github.com/Henry-Sarabia/igdb/v2"
)

var (
	Client = config.InitClient()
)

func GetGamesDataByName(name string, limit int) []*igdb.Game {
	games, err := Client.Games.Search(
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
	games, err := Client.Games.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return games
}

func GetCharacterDataByName(name string, limit int) []*igdb.Character {
	characters, err := Client.Characters.Search(
		name,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return characters
}

func GetGenresDataByIDs(ids []int, limit int) []*igdb.Genre {
	genres, err := Client.Genres.List(
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
	themes, err := Client.Themes.List(
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
	platforms, err := Client.Platforms.List(
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
	involvedCompanies, err := Client.InvolvedCompanies.List(
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
	involvedCompanies, err := Client.Companies.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return involvedCompanies
}

func GetCoversDataByIDs(id int, limit int) *igdb.Cover {
	covers, err := Client.Covers.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return covers
}

func GetMugShotsDataByIDs(id int, limit int) *igdb.CharacterMugshot {
	Mugshots, err := Client.CharacterMugshots.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return Mugshots
}

func GetLogosDataByIDs(id int, limit int) *igdb.CompanyLogo {
	logos, err := Client.CompanyLogos.Get(
		id,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	if err != nil {
		log.Fatal(err)
	}

	return logos
}

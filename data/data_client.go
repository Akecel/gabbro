package data

import (
	"github.com/akecel/gabbro/config"

	"github.com/Henry-Sarabia/igdb/v2"
)

var (
	Client = config.InitClient()
)

func GetGamesDataByName(name string, limit int) ([]*igdb.Game, error) {
	games, err := Client.Games.Search(
		name,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return games, err
}

func GetGamesDataByIDs(ids []int, limit int) ([]*igdb.Game, error) {
	games, err := Client.Games.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return games, err
}

func GetCharacterDataByName(name string, limit int) ([]*igdb.Character, error) {
	characters, err := Client.Characters.Search(
		name,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return characters, err
}

func GetGenresDataByIDs(ids []int, limit int) ([]*igdb.Genre, error) {
	genres, err := Client.Genres.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return genres, err
}

func GetThemesDataByIDs(ids []int, limit int) ([]*igdb.Theme, error) {
	themes, err := Client.Themes.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return themes, err
}

func GetPlatformsDataByIDs(ids []int, limit int) ([]*igdb.Platform, error) {
	platforms, err := Client.Platforms.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return platforms, err
}

func GetInvolvedCompaniesDataByIDs(ids []int, limit int) ([]*igdb.InvolvedCompany, error) {
	involvedCompanies, err := Client.InvolvedCompanies.List(
		ids,
		igdb.SetFields("*"),
		igdb.SetLimit(limit),
	)

	return involvedCompanies, err
}

func GetCompaniesDataByID(id int) (*igdb.Company, error) {
	involvedCompanies, err := Client.Companies.Get(
		id,
		igdb.SetFields("*"),
	)

	return involvedCompanies, err
}

func GetCoversDataByID(id int) (*igdb.Cover, error) {
	covers, err := Client.Covers.Get(
		id,
		igdb.SetFields("*"),
	)

	return covers, err
}

func GetMugShotsDataByID(id int) (*igdb.CharacterMugshot, error) {
	mugshots, err := Client.CharacterMugshots.Get(
		id,
		igdb.SetFields("*"),
	)

	return mugshots, err
}

func GetLogosDataByID(id int) (*igdb.CompanyLogo, error) {
	logos, err := Client.CompanyLogos.Get(
		id,
		igdb.SetFields("*"),
	)

	return logos, err
}

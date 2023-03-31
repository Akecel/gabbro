package tests

import (
	"testing"

	"github.com/akecel/gabbro/responses"
)

func TestGameResponse(t *testing.T) {
	game := responses.GameResponse{
		Name:        "Outer Wilds",
		Description: "Outer Wilds is a critically-acclaimed and award-winning open world mystery about a solar system trapped in an endless time loop.",
		Genres:      "Puzzle, Simulator, Adventure, Indie",
		Themes:      "Science fiction, Action, Open world, Mystery",
		DLCs:        "Outer Wilds: Echoes of the Eye",
		Companies:   "Annapurna Interactive, Mobius Digital",
		Platforms:   "PC (Microsoft Windows), Nintendo Switch, PlayStation 4, PlayStation 5, Xbox Series X|S, Xbox One",
		ReleaseDate: "2019-05-28",
		URL:         "https://www.igdb.com/games/outer-wilds",
		Rating:      9.2,
	}

	if game.Name != "Outer Wilds" {
		t.Errorf("Incorrect Name field of GameResponse.")
	}
	if game.Description != "Outer Wilds is a critically-acclaimed and award-winning open world mystery about a solar system trapped in an endless time loop." {
		t.Errorf("Incorrect Description field of GameResponse")
	}
	if game.Genres != "Puzzle, Simulator, Adventure, Indie" {
		t.Errorf("Incorrect Genres field of GameResponse.")
	}
	if game.Themes != "Science fiction, Action, Open world, Mystery" {
		t.Errorf("Incorrect Themes field of GameResponse.")
	}
	if game.DLCs != "Outer Wilds: Echoes of the Eye" {
		t.Errorf("Incorrect DLCs field of GameResponse.")
	}
	if game.Companies != "Annapurna Interactive, Mobius Digital" {
		t.Errorf("Incorrect Companies field of GameResponse.")
	}
	if game.Platforms != "PC (Microsoft Windows), Nintendo Switch, PlayStation 4, PlayStation 5, Xbox Series X|S, Xbox One" {
		t.Errorf("Incorrect Platforms field of GameResponse.")
	}
	if game.ReleaseDate != "2019-05-28" {
		t.Errorf("Incorrect ReleaseDate field of GameResponse.")
	}
	if game.URL != "https://www.igdb.com/games/outer-wilds" {
		t.Errorf("Incorrect URL field of GameResponse.")
	}
	if game.Rating != 9.2 {
		t.Errorf("Incorrect Rating field of GameResponse.")
	}
}

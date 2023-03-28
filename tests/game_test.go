package tests

import (
	"testing"

	"github.com/akecel/gabbro/config"
	"github.com/akecel/gabbro/cmd"
)

func TestGetGameByNameCmd(t *testing.T) {
	config.InitTestConfig()
	command := cmd.NewGameCmd()

	value := []string {"Outer Wilds"}
	cmd.GetGame(command, value)
}
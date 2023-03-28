package tests

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/akecel/gabbro/cmd"
)

func TestGetGameByNameCmd(t *testing.T) {
	command := cmd.NewGetGameByNameCmd()
	buf := new(bytes.Buffer)
	command.SetOut(buf)

	err := command.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
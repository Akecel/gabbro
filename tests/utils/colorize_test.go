package tests

import (
	"testing"

	"github.com/fatih/color"

	"github.com/akecel/gabbro/utils"
)

func TestSetColor(t *testing.T) {
	// Appel de la fonction SetColor
	setColorFunc := utils.SetColor()

	// Vérification que la fonction retournée est une fonction valide
	if setColorFunc == nil {
		t.Errorf("SetColor() n'a pas retourné une fonction valide.")
	}

	// Vérification que la fonction retournée renvoie la chaîne de caractères attendue
	expected := color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()("test")
	if setColorFunc("test") != expected {
		t.Errorf("SetColor() a retourné la chaîne de caractères incorrecte.")
	}
}

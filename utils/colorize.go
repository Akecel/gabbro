package utils

import (
	"github.com/fatih/color"
)

func SetColor() (func(a ...interface{}) string) {
	return color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
}
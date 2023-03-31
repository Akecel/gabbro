package utils

import (
	"math"
)

func ReadableFloatNb(nb float64) float64 {
	return math.Round(nb/10*100) / 10
}
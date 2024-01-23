package app

import (
	"fmt"
	"math"
)

func Rounding(n float64) float64 {
	multipled := n * 10.0

	rounded := math.Round(multipled)

	result := rounded / 10.0
	fmt.Printf("%.2f\n", result)
	return result
}

package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempsMap := make(map[int][]float64)

	for _, temperature := range temperatures {
		key := int(math.Floor(temperature/10)) * 10
		tempsMap[key] = append(tempsMap[key], temperature)
	}

	for k, v := range tempsMap {
		fmt.Printf("%d: %v\n", k, v)
	}
}
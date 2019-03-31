package main

import (
	"math"
	"math/rand"
)

func GenInt(i float64) int {
	return rand.Intn(int(math.Pow(10, i)))
}

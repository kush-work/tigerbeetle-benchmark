package utils

import (
	"math"
	"math/rand"
	"time"
)

func RandomNumberGenerator() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1
}

func GetAmountInLowestForm(amount float64) uint {
	// Multiply by 100 to convert dollars to cents
	cents := amount * 100

	// Round cents to the nearest whole number
	roundedCents := math.Round(cents)

	// Check if rounded cents exceeds the maximum value for uint
	if roundedCents > math.MaxUint {
		return uint(math.MaxUint)
	}

	// Convert the cents to uint and return
	return uint(roundedCents)
}

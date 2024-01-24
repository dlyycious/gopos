package randomgoidpackage

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(length int) int {
	// Set the seed based on the current time
	rand.Seed(time.Now().UnixNano())

	// Calculate the minimum and maximum values based on the desired length
	min := intPow(10, length-1)
	max := intPow(10, length) - 1

	// Generate a random number within the specified range
	return rand.Intn(max-min+1) + min
}

// Helper function to calculate the power of an integer
func intPow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

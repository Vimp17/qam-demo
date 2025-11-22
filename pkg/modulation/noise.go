package modulation

import (
	"math/rand"
)

// AddGaussianNoise adds AWGN with given sigma
func AddGaussianNoise(symbols []Complex, sigma float64, rng *rand.Rand) []Complex {
	noisy := make([]Complex, len(symbols))
	for i, s := range symbols {
		noisy[i] = Complex{
			I: s.I + rng.NormFloat64()*sigma,
			Q: s.Q + rng.NormFloat64()*sigma,
		}
	}
	return noisy
}

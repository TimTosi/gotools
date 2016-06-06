package random

import "math/rand"

// -----------------------------------------------------------------------------

// GenerateBytes returns a random `[]byte` of size `len`.
//
// NOTE: it's obvious but do not use for cryptography purposes.
func (rg *RandGenerator) GenerateBytes(len int) []byte {
	b := make([]byte, len)

	for j := 0; j < len; {
		rInt := rg.Int63()
		for i := 0; i < 8 && j < len; i++ {
			b[j] = byte(rInt >> 8)
			rInt = rInt >> 8
			j++
		}
	}
	return b
}

// -----------------------------------------------------------------------------

// RandGenerator is a structure that generates pseudo-random values from
// a `rand.Source`.
type RandGenerator struct {
	rand.Source
}

// NewRandomGenerator returns a new `*RandGenerator`.
func NewRandomGenerator(seed int64) *RandGenerator {
	return &RandGenerator{Source: rand.NewSource(seed)}
}

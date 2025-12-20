package main

import (
	"crypto/rand"
	"encoding/binary"
)

var rngSeed uint64

func init() {
	// Initialize seed with cryptographically secure random bytes
	b := make([]byte, 8)
	rand.Read(b)
	rngSeed = binary.BigEndian.Uint64(b)
}

func GenerateRandomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[randIntn(len(letters))]
	}
	return string(b)
}

// randIntn returns, as an int, a non-negative pseudo-random number in [0,n).
// Uses xorshift64* algorithm with package-level seed that persists between calls.
func randIntn(n int) int {
	// xorshift64* algorithm
	rngSeed ^= rngSeed >> 12
	rngSeed ^= rngSeed << 25
	rngSeed ^= rngSeed >> 27
	rngSeed *= 0x2545F4914F6CDD1D
	
	result := int(rngSeed % uint64(n))
	if result < 0 {
		result = -result
	}
	return result
}

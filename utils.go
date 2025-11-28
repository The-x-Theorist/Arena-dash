package main

func GenerateRoomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[randIntn(len(letters))]
	}
	return string(b)
}

// randIntn returns, as an int, a non-negative pseudo-random number in [0,n).
func randIntn(n int) int {
	// simple xorshift, fallback if no import available
	// using package-level seed
	var seed = int64(0xabc12345)
	seed ^= seed << 13
	seed ^= seed >> 17
	seed ^= seed << 5
	if seed < 0 {
		seed = -seed
	}
	return int(seed % int64(n))
}

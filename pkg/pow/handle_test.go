package pow

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkVerifyPoW(b *testing.B) {
	nonce := rand.Intn(math.MaxInt64)
	for i := 0; i < b.N; i++ {
		salt, _ := generateSalt()
		_, _ = verifyPoW(salt, nonce)
	}
}

func BenchmarkFindNonce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		salt, _ := generateSalt()
		_, _ = findNonce(salt)
	}
}

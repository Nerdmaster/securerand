package securerand

import (
	"crypto/rand"
	"math"
	"math/big"
)

var maxInt64 = new(big.Int).SetInt64(math.MaxInt64)
var maxUint64 = new(big.Int).SetUint64(math.MaxUint64)

// Source is a cryptographically secure random number source.  It holds no
// internal state, as the rand.Source implementation merely calls crypto/rand
// methods, and seeding is unused.
type Source struct{}

// Seed implements rand.Source, but panics if called, as cryptographically
// secure randomness obviously can't be seeded
func (s Source) Seed(seed int64) {
	panic("securerand.Source cannot be seeded")
}

// Int63 implements rand.Source, returning a cryptographically secure number
// from 0 to 1<<63
func (s Source) Int63() int64 {
	return rnd(maxInt64).Int64()
}

// Uint64 implements rand.Source64
func (s Source) Uint64() uint64 {
	return rnd(maxUint64).Uint64()
}

// rnd internalizes the common logic of pulling crypto-safe data and returning it
func rnd(*big.Int) *big.Int {
	var val, err = rand.Int(rand.Reader, maxUint64)
	// If we couldn't read from /dev/urandom, it's safe to panic.  Something very
	// bad is going on.
	if err != nil {
		panic(err)
	}

	return val
}

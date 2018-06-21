package cryptosource

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

var maxInt64 = new(big.Int).SetInt64(math.MaxInt64)
var maxUint64 = new(big.Int).SetUint64(math.MaxUint64)

// Source is a cryptographically secure random number source.  It holds no
// internal state, as the rand.Source implementation merely calls crypto/rand
// methods, and seeding is unused.
type Source struct{}

// Seed implements rand.Source, but does nothing, as cryptographically secure
// randomness obviously can't be seeded
func (s Source) Seed(seed int64) {}

// Int63 implements rand.Source, returning a cryptographically secure number
// from 0 to 1<<63
func (s Source) Int63() int64 {
	var val, err = crand.Int(crand.Reader, maxInt64)
	// Suppress errors by returning a PRNG value - this shouldn't happen unless
	// somehow /dev/urandom is broken, in which case we have much bigger problems
	if err != nil {
		return rand.Int63()
	}
	return val.Int64()
}

// Uint64 implements rand.Source64
func (s Source) Uint64() uint64 {
	var val, err = crand.Int(crand.Reader, maxUint64)
	// Suppress errors by returning a PRNG value - this shouldn't happen unless
	// somehow /dev/urandom is broken, in which case we have much bigger problems
	if err != nil {
		return rand.Uint64()
	}
	return val.Uint64()
}

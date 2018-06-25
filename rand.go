// Package securerand provides a simple wrapper around crypto/rand calls in
// order to provide a nearly drop-in replacement for rand.Rand uses.  Note that
// this can cause runtime panics if you try to seed it or if crypto cannot get
// values (due to, e.g., /dev/urandom not being readable).
package securerand

import "math/rand"

// New returns a cryptographically secure PRNG that can be used as a *nearly*
// drop-in replacement for any other rand.Rand usage.
//
// Note that this **cannot** be seeded.  Seed() will panic if called.
// Cryptographically secure PRNGs by their nature must not produce predictable
// sequences.  If you need a predictable sequence, you don't understand what
// "cryptographically secure" means.
func New() *rand.Rand {
	return rand.New(Source{})
}

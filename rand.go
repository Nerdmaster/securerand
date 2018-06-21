package securerand

import "math/rand"

// New returns a cryptographically secure PRNG that can be used as a *nearly*
// drop-in replacement for any other rand.Rand usage.
//
// Note that this **cannot** be seeded.  The Seed() function does nothing.
// Cryptographically secure PRNGs by their nature must not produce predictable
// sequences.  If you need a predictable sequence, you don't understand what
// "cryptographically secure" means.
func New() *rand.Rand {
	return rand.New(Source{})
}

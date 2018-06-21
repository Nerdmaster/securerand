package securerand_test

import (
	"fmt"
	"math/rand"

	"github.com/Nerdmaster/securerand"
)

// This is a dummy assignment that ensures we conform to the Source64
// interface.  For some reason rand.Rand doesn't actually offer any way to
// validate this through standard code, so if this compiles, we know we're
// good.
var _ = rand.Source64(securerand.Source{})

// This simple example shows how you might create a random number generator
// while also proving that seeding is a bad idea
func Example() {
	var r = securerand.New()

	defer func() {
		var r = recover()
		if r == nil {
			fmt.Println("Seed() should have panicked, but didn't!")
		}

		fmt.Println("Panic message: " + r.(string))
	}()

	r.Seed(1)

	// Output:
	// Panic message: securerand.Source cannot be seeded
}

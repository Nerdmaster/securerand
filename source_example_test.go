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
// while also proving that seeding doesn't do anything
func Example() {
	var r = securerand.New()
	r.Seed(1)
	var first = r.Int()
	r.Seed(1)
	var second = r.Int()

	fmt.Println(first == second)
	// Output:
	// false
}

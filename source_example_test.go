package securerand_test

import (
	"fmt"

	"github.com/Nerdmaster/securerand"
)

// This simple example shows how you might create a random number generator
// that could be used exactly the same as if you'd simply imported "math/rand".
// This also shows that seeding is a bad idea.
func Example() {
	var rand = securerand.New()

	defer func() {
		var rec = recover()
		if rec == nil {
			fmt.Println("Seed() should have panicked, but didn't!")
		}

		fmt.Println("Panic message: " + rec.(string))
	}()

	// Roll a die as if we'd imported math/rand
	var _ = rand.Intn(6) + 1
	fmt.Println("Rolled die successfully")

	rand.Seed(1)
	// Output:
	// Rolled die successfully
	// Panic message: securerand.Source cannot be seeded
}

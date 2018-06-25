package securerand_test

import (
	"fmt"
	"math/rand"
	"testing"

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

func TestRange(t *testing.T) {
	var max int64 = 1000
	var x int64

	var r = securerand.New()

	// This is far from perfect, but it gives us a decent shot at catching problems
	for x = 0; x < max; x++ {
		var rndVal = r.Int63n(max)
		if rndVal >= max {
			t.Errorf("r.Int63n(%d) got too high a value: %d", max, rndVal)
			break
		}
	}

	for x = 0; x < max; x++ {
		var rndVal = r.Int63n(max)
		if rndVal < 0 {
			t.Errorf("r.Int63n(%d) got a negative value: %d", max, rndVal)
			break
		}
	}
}

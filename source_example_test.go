package cryptosource_test

import (
	"fmt"

	"github.com/Nerdmaster/cryptosource"
)

// This simple example shows how you might create a random number generator
// while also proving that seeding doesn't do anything
func Example() {
	var r = cryptosource.New()
	r.Seed(1)
	var first = r.Int()
	r.Seed(1)
	var second = r.Int()

	fmt.Println(first == second)
	// Output:
	// false
}

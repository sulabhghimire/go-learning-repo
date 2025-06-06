package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Generating some random number

	// Auto seeding --> returns [0, n) --> from [0, 101) i.e [0, 100]
	fmt.Println("Between 0 and 100", rand.Intn(101))

	// Between [1, 100]
	fmt.Println("Between 1 and 100", rand.Intn(100)+1)

	// Seed a random number generater and generate same set of random numbers
	val := rand.New(rand.NewSource(42))
	fmt.Println("Random number with fix seed:", val.Intn(101))

	// Using unix time as seed
	val = rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Random number with unix seed:", val.Intn(101))

	// Random Floating Numbers
	fmt.Println("Rnadom Floating Numbers:", rand.Float64()) // [0.0 to 1.0)
}

/*
Pseudo-Random Number Generation (PRNG)
a. Seed: starting point for generating a squeqnce of random numbers, setting same seed we can generate
same set of random numbers each time we run our program. Generally Current time is used as a new seed
b. rand.Intn(n)
c. rand.Float64()

Considerations
a. Deterministic Nature
b. Thread Safety
c. Cryptographic Safety
*/

package go_koans

import "fmt"

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == "apple") // slices seem like arrays
	assert(len(fruits) == 3)  // in nearly all respects

	tasty_fruits := fruits[1:3]           // we can even slice slices
	assert(tasty_fruits[0] == "orange") // slices of slices also share the underlying data

	pregnancy_slots := []string{"baby", "baby", "lemon"}
	// fmt.Printf("%d", cap(pregnancy_slots))
	assert(cap(pregnancy_slots) == 3) // the capacity is initially the length

	pregnancy_slots = append(pregnancy_slots, "baby!")
	assert(len(pregnancy_slots) == 4) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancy_slots) == 4) // but with better optimizations

	pregnancy_slots = append(pregnancy_slots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancy_slots) == __int__) // append() can take N arguments to append to the slice
	assert(cap(pregnancy_slots) == __int__) // the capacity optimizations have a guessable algorithm
}

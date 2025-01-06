package go_koans

import "fmt"

type coolNumber int

type babiGila coolNumber

func (cn coolNumber) multiplyByTwo() int {
	uhui := babiGila(cn)
	fmt.Print(uhui)

	return int(cn) * 2
}

func aboutTypes() {
	i := coolNumber(4)
	assert(i == coolNumber(4))     // values can be converted between compatible types
	assert(i.multiplyByTwo() == 8) // you can add methods on any type you define
}

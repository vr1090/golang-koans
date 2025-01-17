package main

import (
	"fmt"
)

func main() {
	var i int = 20
	var f float32 = float32(i)

	fmt.Println(i)
	fmt.Println(f)

	const value = 20
	i = value
	f = value

	fmt.Println(i)
	fmt.Println(f)

}

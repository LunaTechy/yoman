package main

import (
	"fmt"
)

func clo() func() int {
	i := 5
	return func() int {
		i++
		return (i)
	}
}
func main() {
	fmt.Println("hey closure")
	s := clo()
	fmt.Println(s())
	inf := func(i int) {
		fmt.Println(i)
		 fmt.Println((5 + 71))
	}
	inf(4)
}

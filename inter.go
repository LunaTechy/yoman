package main

import (
	"fmt"
	"strings"
)

type inte interface {
	up()
	low()
}
type strin struct {
	s string
}

func (is *strin) up() {
	fmt.Println(strings.ToUpper(is.s))
}

func (is *strin) low() {
	fmt.Println(strings.ToLower(is.s))
}

func main() {

	si := strin{"Hello"}
	si.up()
	si.low()

}

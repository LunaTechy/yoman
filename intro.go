package main

import "fmt"

type q struct {
	name string
	age  int
}

func main() {

	var a int
	a = 5
	s := 5
	fmt.Println(a, s)

	var f [5]int
	f[1] = 4
	fmt.Println(f)

	g := f[:]
	fmt.Println(g)

	g = append(g, 1, 2)
	g = append(g, f[:]...)
	fmt.Println(g)

	const y = 5
	fmt.Println(y)

	fq := q{"ji", 3}
	fmt.Println(fq)

	fmt.Println(fq.name)

	//gg := []byte{"hry"}
	//fmt.Print(gg)

	gg := "hey all"
	ggf := []byte(gg)
	fmt.Printf("%s", ggf)
}

package main

import (
	"fmt"
	"sync"
)

func add(){
	for i := 5, i <=20 ; i++{
		fmt.Println(i)
	}
}

func main(){
	
	go add()
	add()
}
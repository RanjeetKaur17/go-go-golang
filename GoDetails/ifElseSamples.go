package main

import (
	"fmt"
)

func runIfElseSamples(){

	fmt.Println("-----------If/Else Samples----------")

	val := 0
	a, b, c := 4, 5, 30

	if i := a * b; i < c {
		val += i
	}
	fmt.Println(val)

	if a < b {
		fmt.Println(1)
	} else if a > b {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}

	fmt.Println("------------------------------------")

}


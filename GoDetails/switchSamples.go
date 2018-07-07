package main

import (
	"fmt"
)

func runSwitchSamples(){

	fmt.Println("------------Switch Samples----------")

	a, b := 1, 2
	switch a {
		case 1: fmt.Println("a")
		default: fmt.Println("NA")
	}

	switch {
		case a>b:fmt.Println(1)
		case b>a:fmt.Println(-1)
		default: fmt.Println(0)
	}

	fmt.Println("------------------------------------")

}


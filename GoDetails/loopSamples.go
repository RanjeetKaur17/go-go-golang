package main

import (
	"fmt"
)

func runLoopSamples(){

	fmt.Println("-------------Loop Samples-----------")

	val := 0
	for i := 0; i < 10; i++ {
		val += i
	}
	fmt.Println(val)

	i := 0
	val = 0
	for i < 10 {
		val += i
		i++
	}
	fmt.Println(val)

	i = 0
	val = 0
	for {
		val += i
		i++;
		if i == 10 {
			break;
		}
	}

	fmt.Println(val)

	fmt.Println("------------------------------------")

}


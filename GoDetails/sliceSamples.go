package main

import (
	"fmt"
	"strconv"
)

func runSliceSamples(){

	fmt.Println("------------Slice Samples-----------")

	var sl1 = []int{}
	fmt.Println(sl1)

	var sl2 = []int{1,2,3}
	fmt.Println(sl2)

	var sl3 = make([]int, 5, 5)
	fmt.Println(sl3)

	var sl4 = make([]int, 5)
	fmt.Println(sl4)

	length := len(sl1)
	fmt.Println(length)

	capacity := cap(sl1)
	fmt.Println(capacity)

	for i, val := range sl2 {
		fmt.Println(strconv.FormatInt(int64(i), 10) + " -> " + strconv.FormatInt(int64(val), 10))
	}

	sl2 = append(sl2, 4)

	for i := 0; i < len(sl2); i++ {
		fmt.Println(strconv.FormatInt(int64(i), 10) + " -> " + strconv.FormatInt(int64(sl2[i]), 10))
	}

	fmt.Println("------------------------------------")

}


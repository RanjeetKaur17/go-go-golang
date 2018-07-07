package main

import (
	"fmt"
	"strconv"
)

func runArraySamples(){

	fmt.Println("------------Array Samples-----------")

	var arr1 = [5]int{}
	fmt.Println(arr1)

	var arr2 = [...]int{1,2,3}
	fmt.Println(arr2)

	length := len(arr1)
	fmt.Println(length)

	for i, val := range arr2 {
		fmt.Println(strconv.FormatInt(int64(i), 10) + " -> " + strconv.FormatInt(int64(val), 10))
	}

	fmt.Println("------------------------------------")

}


package main

import (
	"fmt"
	"strconv"
)

func runMapSamples(){

	fmt.Println("-------------Map Samples------------")

	map1 := make(map[int]string)
	fmt.Println(map1)

	map2 := map[int]string{}
	fmt.Println(map2)

	map3 := map[int]string{1:"A", 2:"B"}
	fmt.Println(map3)

	key := 1
	map1[key] = "Test"
	value := map1[key]
	fmt.Println(strconv.FormatInt(int64(key), 10) + " -> " + value)

	if _, ok := map1[key];ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exist")
	}

	delete(map1, key)

	if _, ok := map1[key];ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exist")
	}

	map1[key] = "Test2"

	n := len(map1)
	fmt.Println(n)

	for key, value := range map1 {
		fmt.Println(strconv.FormatInt(int64(key), 10) + " -> " + value)
	}



	fmt.Println("------------------------------------")

}


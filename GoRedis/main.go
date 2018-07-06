package main

import "go-go-golang/GoRedis/cache"

func main() {

	cache.InitializeRedis()

	println("Setting Testkey -> TestValue")
	cache.SetValue("TestKey", "TestValue")

	println("Getting TestKey")
	value, err := cache.GetValue("TestKey")

	if err == nil {
		println("Value Returned : " + value.(string))
	} else {
		println("Getting Value Failed with error : " + err.Error())
	}

}
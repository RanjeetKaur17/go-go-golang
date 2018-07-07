package main

import "fmt"

type Person struct{
	Id int
	Name string
	Age int
	Email string
}

func runStructSamples(){

	fmt.Println("----------Struct Samples-----------")

	//Object with No arguments
	person1 := Person{}
	fmt.Println(person1)

	//Object with all arguments
	person2 := Person{Id:1,Name:"JaneDoe",Age:24,Email:"x@abc.com"}
	fmt.Println(person2)

	//Object with partial arguments
	person3 := Person{Name:"JohnDoe"}
	fmt.Println(person3)

	fmt.Println("------------------------------------")
}

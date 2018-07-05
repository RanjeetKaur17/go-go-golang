package main

import (
	"fmt"
	"go-go-golang/BuildingBlocks/importSamples/default"
	NP "go-go-golang/BuildingBlocks/importSamples/namedPath"
	. "go-go-golang/BuildingBlocks/importSamples/unNamedPath"
	_ "go-go-golang/BuildingBlocks/importSamples/ignoredPath"
)

func main(){

	importSamples()

	variablesSamples()

	structSamples()

	functionSamples()

}

func importSamples() {
	fmt.Print("--------Printing Import Samples-------\n")

	//Default Import
	fmt.Print(defaultPath.Foo1() + "\n")

	//Named Import
	fmt.Print(NP.Foo3() + "\n")

	//Unamed Import
	fmt.Print(Foo4() + "\n")

	fmt.Print("--------------------------------------\n")
}

func variablesSamples() {

	fmt.Print("-------Printing Variable Samples-------\n")

	//Variables
	var name1 string = "JohnDoe1"
	fmt.Print(name1 + "\n")

	var name2 = "JohnDoe2"
	fmt.Print(name2 + "\n")

	name3 := "JohnDoe3"
	fmt.Print(name3 + "\n")

	//Constants
	const NAME = "NAME"
	fmt.Print(NAME + "\n")

	const (
		FIRST_NAME = "FIRSTNAME"
		LAST_NAME = "LASTNAME"
	)

	fmt.Print(FIRST_NAME + "-" + LAST_NAME + "\n")

	fmt.Print("--------------------------------------\n")
}

func structSamples() {
	fmt.Print("---------Printing Struct Samples-------\n")

	var name = Name{FirstName:"John", LastName:"Doe"}
	fmt.Print(name.FirstName + name.LastName + "\n")

	fmt.Print("--------------------------------------\n")
}

func functionSamples() {

	fmt.Print("-------Printing Function Samples-------\n")

	//Call by Value
	formattedName, formattedAge := format("JohnDoe", 25)
	fmt.Print(formattedName + "\n")
	fmt.Print(formattedAge + "\n")

	//Call by Pointer
	name := "JohnDoe"
	formatName(&name)
	fmt.Print(name + "\n")

	fmt.Print("--------------------------------------\n")
}




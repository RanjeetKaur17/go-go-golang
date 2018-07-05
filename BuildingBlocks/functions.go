package main

import "strconv"

func format(name string, age int64) (string, string) {
	return "Name : " + name, "Age : " + strconv.FormatInt(age, 10)
}

func formatAge(age int64) (string) {
	return "Age : " + strconv.FormatInt(age, 10)
}

func formatName(name *string) {
	*name = "Name : " + *name
}
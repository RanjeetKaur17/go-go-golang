package test

import (
	"testing"
	"go-go-golang/GoTesting/calculator"
)

func TestAdd(t *testing.T) {

	sum := calculator.Add(1,2)

	if sum != 3 {
		t.Error("Sum Did Not Match!!")
		t.Fail()
	} else {
		t.Log("Add Test Passed")
	}
}

func TestSub(t *testing.T) {

	sum := calculator.Sub(1,2)

	if sum != -1 {
		t.Error("Diff Did Not Match!!")
		t.Fail()
	} else {
		t.Log("Add Test Passed")
	}
}

func TestMul(t *testing.T) {

	sum := calculator.Mul(1,2)

	if sum != 2 {
		t.Error("Product Did Not Match!!")
		t.Fail()
	} else {
		t.Log("Add Test Passed")
	}
}

func TestDiv(t *testing.T) {

	sum := calculator.Div(1,2)

	if sum != 0 {
		t.Error("Quotient Did Not Match!!")
		t.Fail()
	} else {
		t.Log("Add Test Passed")
	}
}
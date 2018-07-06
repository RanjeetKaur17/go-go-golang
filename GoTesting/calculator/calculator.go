package calculator

func Add(a int,b int) int {
	return a+b;
}

func Mul(a int,b int) int {
	return a*b;
}

func Sub(a int,b int) int {
	return a-b;
}

func Div(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a/b
}
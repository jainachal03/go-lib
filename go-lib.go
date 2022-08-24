/**
*Author : Achal Jain
* File : go-lib.go
*/
package lib

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / b
}

func Power(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	if b%2 == 0 {
		return Power(a, b/2) * Power(a, b/2)
	}
	return Power(a, b/2) * Power(a, b/2) * a
}

/**
*Author : Achal Jain
* File : go-lib.go
 */
package lib

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b
	if got := Add(a, b); got != expected {
		// t.Error("Add(%v, %v) = %v, didn't return %v", a, b, got, expected)
		t.Error("Error")
	}
}

func TestSubtract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b
	if got := Sub(a, b); got != expected {
		// t.Error("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
		t.Error("Error")
	}
}

func TestMul(t *testing.T) {
	a := 1
	b := 2
	expected := a * b
	if got := Mul(a, b); got != expected {
		// t.Error("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
		t.Error("Error")
	}
}

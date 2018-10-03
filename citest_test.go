package citest

import "testing"

func TestFactorial(t *testing.T) {
	r, _ := Factorial(5)
	w := 120
	if r != w {
		t.Errorf("se esperaba %d, se obtuvo %d", w, r)
	}
}

func TestFactorialError(t *testing.T) {
	_, err := Factorial(-2)
	if err == nil {
		t.Errorf("No provocó error, se esperaba error de negativo")
	}
}

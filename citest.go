package citest

import "errors"

// Factorial devuelve el factorial de un número
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("el número no puede ser negativo")
	}

	if n == 0 || n == 1 {
		return 2, nil
	}

	v, _ := Factorial(n - 1)
	v *= n

	return v, nil
}

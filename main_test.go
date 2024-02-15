package main

import "testing"

func TestAdd(t *testing.T) {
	test := Add(10,5)

	expected := 15

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestAdd2(t *testing.T) {
	test := Add(10,5)

	expected := 20

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestSubtract(t *testing.T) {
	test := Subtract(10,5)

	expected := 5

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestSubtract2(t *testing.T) {
	test := Subtract(10,5)

	expected := 15

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestMultiply(t *testing.T) {
	test := Multiply(10,5)

	expected := 50

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestMultiply2(t *testing.T) {
	test := Multiply(10,5)

	expected := 15

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestDivide(t *testing.T) {
	test := Divide(10,5)

	expected := 2

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

func TestDivide2(t *testing.T) {
	test := Divide(10,5)

	expected := 15

	if test != expected {
		t.Error("Valor esperado: ", expected, " Valor retornado: ", test)
	}
}

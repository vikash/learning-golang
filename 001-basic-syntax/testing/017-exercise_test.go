package main

import "testing"

func TestMultiply(t *testing.T) {
	a, b := 2, 4
	expectedOutput := 8

	output := multiply(a, b)
	if output != expectedOutput {
		t.Errorf("Expected %v\tGot %v", expectedOutput, output)
	}
}

func TestHello(t *testing.T) {
	name := "Joe"
	expectedOutput := "hello Joe"

	output := hello(name)
	if output != expectedOutput {
		t.Errorf("Expected %v\tGot %v", expectedOutput, output)
	}
}

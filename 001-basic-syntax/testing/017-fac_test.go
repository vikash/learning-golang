package main

import "testing"

// go test								Run tests for current package.
// go test ./... 						Tests current and all subpackages.
// go test --bench=. --benchmem			Run benchmarks including memory.

// go test -coverprofile=coverage.out	Output coverage profile in a file.
// go tool cover -func=coverage.out		Function wise coverage from the output file.
// go tool cover -html=coverage.out		Detailed coverage in Browser
func TestFac(t *testing.T) {
	input := 0
	expectedOutput := 1
	output := fac(input)
	if expectedOutput != output {
		t.Fatalf("Factorial %v: Expected %v, Got %v ", input, expectedOutput, output)
	}
}

func TestFacr(t *testing.T) {
	input := 5
	expectedOutput := 120
	output := facr(input)
	if expectedOutput != output {
		t.Fatalf("Factorial %v: Expected %v, Got %v ", input, output, expectedOutput)
	}
}

func BenchmarkFac20(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		fac(20)
	}
}

func BenchmarkFacr20(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		facr(20)
	}
}

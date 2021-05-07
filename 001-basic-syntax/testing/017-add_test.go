package main

import (
	"testing"
)

// go test								Run tests for current package.
// go test ./... 						Tests current and all subpackages.
// go test --bench=. --benchmem			Run benchmarks including memory.
// go test -coverprofile=coverage.out	Output coverage profile in a file.
// go tool cover -func=coverage.out		Function wise coverage from the output file.
// go tool cover -html=coverage.out		Detailed coverage in Browser

func TestAdd(t *testing.T) {
	a, b := 3, 4
	expectedOutput := 7

	output := add(a, b)
	if output != expectedOutput {
		t.Errorf("Expected %v\tGot %v", expectedOutput, output)
	}
}

package functions

import "testing"

// go test								Run tests for current package.
// go test ./... 						Tests current and all subpackages.
// go test --bench=. --benchmem			Run benchmarks including memory.
// go test -coverprofile=coverage.out	Output coverage profile in a file.
// go tool cover -func=coverage.out		Function wise coverage from the output file.
// go tool cover -html=coverage.out		Detailed coverage in Browser

func TestFac(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}

	for _, tc := range tests {
		r := fac(tc.input)
		if tc.output != r {
			t.Fatalf("Factorial %v: Expected %v, Got %v ", tc.input, tc.output, r)
		}
	}
}

func TestFacr(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}

	for _, tc := range tests {
		r := facr(tc.input)
		if tc.output != r {
			t.Fatalf("Factorial %v: Expected %v, Got %v ", tc.input, tc.output, r)
		}
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

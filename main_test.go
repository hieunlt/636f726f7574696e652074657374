package main

import (
	"fmt"
	"testing"
)

// testCase represents a test case for isPrime function
type testCase struct {
	name     string // descriptive name of the test case
	input    int    // input number to test
	expected bool   // expected result
	comment  string // optional comment explaining the test case
}

// runTests executes a slice of test cases with proper error reporting
func runTests(t *testing.T, category string, cases []testCase) {
	for _, tc := range cases {
		testName := fmt.Sprintf("%s/%s(n=%d)", category, tc.name, tc.input)
		t.Run(testName, func(t *testing.T) {
			got := isPrime(tc.input)
			if got != tc.expected {
				t.Errorf("\ninput: %d\nwant: %v\ngot: %v\ncomment: %s",
					tc.input, tc.expected, got, tc.comment)
			}
		})
	}
}

// TestIsPrimeBasicCases tests fundamental properties of prime numbers
func TestIsPrimeBasicCases(t *testing.T) {
	cases := []testCase{
		{"Negative", -7, false, "negative numbers are not prime"},
		{"Zero", 0, false, "zero is not prime"},
		{"One", 1, false, "one is not prime by definition"},
		{"Two", 2, true, "two is the smallest and only even prime"},
		{"Three", 3, true, "three is prime"},
	}
	runTests(t, "Basic", cases)
}

// TestIsPrimeSmallNumbers tests numbers up to 20
func TestIsPrimeSmallNumbers(t *testing.T) {
	cases := []testCase{
		{"Four", 4, false, "2 × 2"},
		{"Five", 5, true, "prime"},
		{"Six", 6, false, "2 × 3"},
		{"Seven", 7, true, "prime"},
		{"Eight", 8, false, "2 × 2 × 2"},
		{"Nine", 9, false, "3 × 3"},
		{"Ten", 10, false, "2 × 5"},
		{"Eleven", 11, true, "prime"},
		{"Twelve", 12, false, "2 × 2 × 3"},
		{"Thirteen", 13, true, "prime"},
		{"Fourteen", 14, false, "2 × 7"},
		{"Fifteen", 15, false, "3 × 5"},
		{"Sixteen", 16, false, "2 × 2 × 2 × 2"},
		{"Seventeen", 17, true, "prime"},
		{"Eighteen", 18, false, "2 × 3 × 3"},
		{"Nineteen", 19, true, "prime"},
		{"Twenty", 20, false, "2 × 2 × 5"},
	}
	runTests(t, "Small", cases)
}

// TestIsPrimeDivisibilityRules tests numbers specifically chosen to test divisibility optimizations
func TestIsPrimeDivisibilityRules(t *testing.T) {
	cases := []testCase{
		{"Multiple of 2", 22, false, "testing even number optimization"},
		{"Multiple of 3", 27, false, "testing divisibility by 3 optimization"},
		{"Multiple of both 2 and 3", 24, false, "testing combined 2 and 3 optimization"},
		{"Prime after multiple of 2", 23, true, "prime after even number"},
		{"Prime after multiple of 3", 31, true, "prime after multiple of 3"},
	}
	runTests(t, "DivisibilityRules", cases)
}

// TestIsPrimePerfectSquares tests perfect squares and their neighbors
func TestIsPrimePerfectSquares(t *testing.T) {
	cases := []testCase{
		{"Perfect Square", 25, false, "5 × 5"},
		{"Before Square", 24, false, "near perfect square"},
		{"After Square", 26, false, "near perfect square"},
		{"Larger Square", 121, false, "11 × 11"},
		{"Prime near Square", 127, true, "prime near perfect square"},
	}
	runTests(t, "PerfectSquares", cases)
}

// TestIsPrimeLargeNumbers tests larger prime and composite numbers
func TestIsPrimeLargeNumbers(t *testing.T) {
	cases := []testCase{
		{"Large Prime", 997, true, "prime"},
		{"Large Composite", 999, false, "3 × 3 × 3 × 37"},
		{"Near Multiple Check", 901, false, "17 × 53"},
		{"Large Prime 2", 7919, true, "prime"},
		{"Special Case", 7917, false, "3 × 2639"},
	}
	runTests(t, "Large", cases)
}

// BenchmarkIsPrime runs performance tests on different number sizes
func BenchmarkIsPrime(b *testing.B) {
	benchCases := []struct {
		name  string
		input int
	}{
		{"Small Prime", 17},
		{"Small Composite", 24},
		{"Medium Prime", 997},
		{"Medium Composite", 1000},
		{"Large Prime", 7919},
		{"Large Composite", 7917},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrime(bc.input)
			}
		})
	}
}

// TestIsPrimeOptimizationBoundary tests numbers around the square root optimization boundary
func TestIsPrimeOptimizationBoundary(t *testing.T) {
	cases := []testCase{
		{"Just Below Sqrt", 99, false, "3 × 3 × 11"},
		{"At Sqrt", 100, false, "2 × 2 × 5 × 5"},
		{"Just Above Sqrt", 101, true, "prime"},
		{"Large Below Sqrt", 7915, false, "5 × 1583"},
		{"Large Above Sqrt", 7920, false, "2 × 2 × 2 × 3 × 3 × 5 × 11"},
	}
	runTests(t, "OptimizationBoundary", cases)
}

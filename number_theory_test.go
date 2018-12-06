package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSumLargestPrimeFactors0(t *testing.T) {
	tests := []struct {
		value int
		sum   int
	}{
		{10, 32},
		{100, 1915},
		{10000, 10118280},
	}
	var result int
	for i := range tests {
		result = numberDivisorList(tests[i].value).sumLargestPrimeFactors(tests[i].value)
		if result != tests[i].sum {
			t.Fatalf("expected %d, received %d\n", tests[i].sum, result)
		}
	}
}

// func TestSumLargestPrimeFactors1(t *testing.T) {
// 	tests := []struct {
// 		value int
// 		sum   int
// 	}{
// 		{10, 32},
// 		{100, 1915},
// 		{10000, 10118280},
// 	}
// 	var result int
// 	for i := range tests {
// 		result = sumLargestPrimeFactors(tests[i].value)
// 		if result != tests[i].sum {
// 			t.Fatalf("expected %d, received %d\n", tests[i].sum, result)
// 		}
// 	}
// }

func BenchmarkSumLargestPrimeFactors0(b *testing.B) {
	for i := 0; i < 6; i++ {
		b.Run(
			fmt.Sprintf("n = 10^%d", i+1),
			func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					numberDivisorList(int(math.Pow10(i + 1))).sumLargestPrimeFactors(int(math.Pow10(i + 1)))
				}
			},
		)
	}
}

func BenchmarkSumLargestPrimeFactors1(b *testing.B) {
	for i := 0; i < 6; i++ {
		b.Run(
			fmt.Sprintf("n = 10^%d", i+1),
			func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					sumLargestPrimeFactors(int(math.Pow10(i + 1)))
				}
			},
		)
	}
}

package main

import (
	"fmt"
	"math"
	"sort"
)

type number struct {
	value     int // Number to be factored
	divisor   int // Smallest divisor of value
	remainder int // Remainder of value / divisor
}
type numbers []*number

func (nums numbers) largestPrimeFactor(n int) int {
	p := nums[n-1].divisor
	for 1 < n {
		if p < nums[n-1].divisor {
			p = nums[n-1].divisor
		}
		n = nums[n-1].remainder
	}
	return p
}

func (nums numbers) factorNumber(n int) map[int]int {
	facts := make(map[int]int)
	for 1 < n {
		facts[nums[n-1].divisor]++
		n = nums[n-1].remainder
	}
	return facts
}

// numberDivisorList returns
func numberDivisorList(n int) numbers {
	facts := make(numbers, n)
	pnums := []int{2, 3}
	for i := range facts {
		facts[i] = &number{
			value:     i + 1,
			divisor:   i + 1, // Assume each i is prime
			remainder: 1,
		}
	}
	var j int
	for 0 < n {
		if n%2 == 0 {
			// n is even
			facts[n-1].divisor = 2
			facts[n-1].remainder = n / 2
		} else {
			// n is odd
			for i := 3; i <= int(math.Sqrt(float64(n))); {
				if n%i == 0 {
					// n is composite
					facts[n-1].divisor = i
					facts[n-1].remainder = n / i
					if facts[n-1].remainder == 1 {
						// Divisor is a prime; insert if new
						if sort.Search(len(pnums), func(index int) bool { return pnums[index] == facts[n-1].divisor }) == len(pnums) {
							pnums = append(pnums, facts[n-1].divisor)
							sort.Ints(pnums)
							fmt.Println("primes:", pnums)
						}
					}
					break
				}
				if pnums[len(pnums)-1] <= i {
					i += 2 // Search for higher prime divisor candidates
					j = 1
				} else {
					i = pnums[j] // Use the next known prime as a divisor
					j++
				}
			}
		}
		n--
	}
	return facts
}

func (nums numbers) Len() int {
	return len(nums)
}

func (nums numbers) Less(i, j int) bool {
	return nums[i].value < nums[j].value
}

func (nums numbers) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// factorList returns an (n+1)x2 dimensional [][]int (f) where
// the ith []int has the following definition: f[i][0] is the
// smallest prime divisor of i and f[i][1] is the remainder of
// i divided by the divisor.
func factorList(n int) [][]int {
	flist := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		flist[i] = make([]int, 2)
		flist[i][0] = i // Assume each i is prime
		flist[i][1] = 1
	}
	for 0 < n {
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				// n is composite
				flist[n][0] = i
				flist[n][1] = n / i
				break
			}
		}
		n--
	}
	return flist
}

// largestPrimeFactor returns the largest prime factor of a given natural number greater than one.
// Source: https://projecteuler.net/problem=642
func largestPrimeFactor(n int) int {
	if n < 2 {
		panic("n must be greater than one to have prime factors")
	}

	pf := factors(n)
	var p int
	for k := range pf {
		if p < k {
			p = k
		}
	}
	return p // pf is guarenteed to have at least one key,value-paired entry, so p will never be less than 2
}

func sumLargestPrimeFactors(n int) int {
	var s int
	for i := 2; i <= n; i++ {
		s += largestPrimeFactor(i)
	}
	return s
}

// factors returns a map of factors of an integer. The keys are the
// factors and the values are the number of times the factor key divides
// the number n. For example, 12 = map[2:2 3:1] = 2^2 * 3.
// Source: https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
func factors(n int) map[int]int {
	if n == 0 {
		panic("cannot factor zero")
	}

	f := make(map[int]int)
	var count int
	for n%2 == 0 {
		count++
		n /= 2
	}
	if 0 < count {
		f[2] = count // n is even
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		count = 0
		for n%i == 0 {
			count++
			n /= i
		}
		if 0 < count {
			f[i] = count // n is divisible by i
		}
	}
	if 2 < n {
		f[n] = 1 // n is prime
	}
	return f
}

// primes returns prime numbers on range [2,n).
func primes(n int) []int {
	// Naive method
	var p []int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			p = append(p, i)
		}
	}
	return p

	// Sieve of Eratosthenes ... TODO
	// pm := make(map[int]bool)
	// for i := 2; i <= n; i++ {
	// 	pm[i] = true
	// }
	// p := 2
	// count := n - 2
	// for p*p <= n {
	// 	if pm[p] {
	// 		for i := p * p; i <= n; i += p {
	// 			pm[i] = false
	// 			count--
	// 		}
	// 	}
	// 	p++
	// }
	// primes := make([]int, 0, count)
	// for k, v := range pm {
	// 	if v {
	// 		primes = append(primes, k)
	// 	}
	// }
	// return primes
}

// isPrime returns true if n is prime and false otherwise.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 && i < n {
			return false // i divides n, but n and i are not equal
		}
	}
	return true
}

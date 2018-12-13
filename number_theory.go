package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
)

type number struct {
	value     int // Number to be factored
	divisor   int // Smallest divisor of value
	nextValue int // Remainder of value / divisor
}

type numbers []*number

// sumLargestPrimeFactors returns the sum of the largest prime factors in
// each number from two up to and including a given number n.
func (nums numbers) sumLargestPrimeFactors(n int) int {
	var s int
	for 1 < n {
		s += nums.largestPrimeFactor(n)
		n--
	}
	return s
}

// largestPrimeFactor returns the largest prime factor of a number in a
// set of numbers.
func (nums numbers) largestPrimeFactor(n int) int {
	p := nums[n-1].divisor
	for 1 < n {
		if p < nums[n-1].divisor {
			p = nums[n-1].divisor
		}
		n = nums[n-1].nextValue
	}
	return p
}

func (nums numbers) factorNumber(n int) map[int]int {
	facts := make(map[int]int)
	for 1 < n {
		facts[nums[n-1].divisor]++
		n = nums[n-1].nextValue
	}
	return facts
}

// numberDivisorList returns a set of numbers from one up to and
// including a given number n.
func numberDivisorList(n int) numbers {
	if n < 1 {
		panic("n must be positive")
	}

	nums := make(numbers, n) // Numbers to return

	// Initialize numbers as 1, 2, ..., n
	for i := range nums {
		nums[i] = &number{
			value:     i + 1,
			divisor:   i + 1, // Initially, assume each i is prime
			nextValue: 1,     // Divisor is indicated as prime if nextValue is one
		}
	}

	// Iterate from n to one finding the smallest divisor and nextValue
	for ; 0 < n; n-- {
		if n%2 == 0 {
			// Case: n is even
			nums[n-1].divisor = 2
			nums[n-1].nextValue = n / 2
			continue
		}

		// Case: n is odd

		// Iterate through known primes if d is less than the
		// current largest prime. Otherwise, iterate
		// from largest prime up to sqrt(n) for each odd divisor
		// candidate.
		// for d <= int(math.Sqrt(float64(n))) {
		for d := 3; d <= n; d += 2 {
			if n%d == 0 {
				// n is composite
				nums[n-1].divisor = d
				nums[n-1].nextValue = n / d
				break
			}
		}
	}
	return nums
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

// primes returns prime numbers on range [2,n). Use eratosthenes instead.
func primes(n int) []int {
	// Naive method
	var p []int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			p = append(p, i)
		}
	}
	return p
}

// eratosthenes returns a list of prime numbers on the range [2,n].
func eratosthenes(n int) []int {
	p := make(map[int]bool) // Indicates if a number is prime (true) or composite (false)
	for i := 2; i <= n; i++ {
		p[i] = true // Initialize all integers as prime
	}

	seive := make([]int, 0, n-2) // Prime numbers to return
	for i := 2; i <= n; i++ {
		if p[i] {
			seive = append(seive, i) // i must be prime at this point
			for j := 2 * i; j <= n; j += i {
				p[j] = false // All multiples of i are not prime
			}
		}
	}
	return seive
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

// importSequence returns a sequence of integers from a csv file.
func importSequence(filename string) ([]int, error) {
	sequence := make([]int, 0, 256)
	data := make([]string, 0, 256)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	var n int
	for i := 0; ; i++ {
		data, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		n, err = strconv.Atoi(data[0])
		if err != nil {
			return nil, err
		}

		sequence = append(sequence, n)
	}
	return sequence, nil
}

// exportSequence writes a seqence of integers to a csv file. Each value will be written on a separate line.

func exportSequence(sequence []int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}

		file, err = os.OpenFile(filename, 0, os.ModePerm)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	s := make([]string, 1)
	for i := range sequence {
		s[0] = strconv.Itoa(sequence[i])
		err = writer.Write(s)
		if err != nil {
			return err
		}

		writer.Flush()
	}

	return writer.Error()
}

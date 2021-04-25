package anagramgame

import (
	"fmt"
	"math/big"
	"unicode"
)

var (
	primeLUT = makePrimeLUT()
)

func makePrimeLUT() map[rune]int {
	primes := makePrimes(26)
	result := make(map[rune]int, 52)
	runes := []rune("abcdefghijklmnopqrstuvwxyz")
	for i, r := range runes {
		result[r] = primes[i]
		result[unicode.ToUpper(r)] = primes[i]
	}
	return result
}

func makePrimes(size int) []int {
	if size == 0 {
		return nil
	}
	primes := []int{2}
	candidate := 3
	for len(primes) < 26 {
		if big.NewInt(int64(candidate)).ProbablyPrime(0) {
			primes = append(primes, candidate)
		}
		candidate += 2
	}
	return primes
}

// IsAnagramPrimes determines if two []runes are an anagram. It panics when
// given non-English letters.
//
// This function maps all of its inputs to prime numbers and compares their
// products to know if the []runes are anagram.
func IsAnagramPrimes(s1, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Result := 1
	s2Result := 1
	for i := range s1 {
		s1Prime, ok := primeLUT[s1[i]]
		if !ok {
			panic(fmt.Sprintf("rune not supported: %s", string([]rune{s1[i]})))
		}
		s1Result = s1Result * s1Prime
		s2Prime, ok := primeLUT[s2[i]]
		if !ok {
			panic(fmt.Sprintf("rune not supported: %s", string([]rune{s2[i]})))
		}
		s2Result = s2Result * s2Prime
	}

	return s1Result == s2Result
}

package anagramgame

import "unicode"

var countLUT = makeCountLUT()

func makeCountLUT() map[rune]int {
	result := make(map[rune]int, 52)
	runes := []rune("abcdefghijklmnopqrstuvwxyz")
	for i, r := range runes {
		result[r] = i
		result[unicode.ToUpper(r)] = i
	}
	return result
}

// IsAnagramCounts determines if two []runes are an anagram. It panics when
// given non-English letters.
//
// This function counts the frequency of all letters in the []runes. If they
// are equal then the arguments are an anagram.
func IsAnagramCounts(s1, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1counts := make([]int, 26)
	s2counts := make([]int, 26)
	for i := range s1 {
		s1counts[countLUT[s1[i]]]++
		s2counts[countLUT[s2[i]]]++
	}
	for i := range s1counts {
		if s1counts[i] != s2counts[i] {
			return false
		}
	}
	return true
}

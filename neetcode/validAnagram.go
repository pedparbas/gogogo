package neetcode

import (
	"slices"
)

/*

	given two strings s and t return true if the two strings are anagrams of each other, otherwise return false

*/

// O(n+m+n)

func isAnagram(s string, t string) bool {
	dict := make(map[rune]int, len(s))
	for _, v := range s {
		dict[v] += 1
	}
	for _, v := range t {
		dict[v] -= 1
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

// O(nlogn + mlogm)
func isAnagramSorting(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes, tRunes := []rune(s), []rune(t)

	slices.Sort(sRunes)
	slices.Sort(tRunes)

	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}

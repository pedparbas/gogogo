package neetcode

import (
	"maps"
	"sort"
	// "strings"
)

/*
	Given an array of strings *strs*, group all anagrams together into sublists.
	You may return the output in any order.

	An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

*/

func groupAnagrams(strs []string) [][]string {
	return [][]string{}
}

func groupAnagramsBForce(strs []string) [][]string {
	// mapped := make(map[string][]string)
	// for _, w := range strs {
	// 	for k, v := range mapped {
	//
	// 	}
	// }

	return [][]string{}
}

// O(n*mlogm)
func groupAnagramsSorting(strs []string) [][]string {
	mapped := make(map[string][]string)
	for _, w := range strs {
		runes := []rune(w)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

		if strlist, ok := mapped[string(runes)]; ok {
			mapped[string(runes)] = append(strlist, w)
		} else {
			mapped[string(runes)] = []string{w}
		}
	}

	var res [][]string
	for _, v := range mapped {
		res = append(res, v)
	}
	return res
}

package array_and_hash

import (
	"sort"
	"strconv"
)

// error passes 117 test case out of 120
func GroupAnagrams(strs *[]string) [][]string {
	dict := make(map[string][]string, len(*strs))

	for _, val := range *strs {
		freq := make([]int, 26)
		for _, ch := range val {
			freq[ch-'a']++
		}
		key := ""
		for _, count := range freq {
			key += strconv.Itoa(count)
		}

		dict[key] = append(dict[key], val)
	}

	ans := make([][]string, 0, len(dict))
	for _, val := range dict {
		ans = append(ans, val)
	}

	return ans
}

// passes all test cases
func GroupAnagrams2(strs []string) [][]string {
	dict := make(map[string][]string, len(strs))

	for _, val := range strs {
		// Convert string to []rune for sorting
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		sortedStr := string(runes)

		dict[sortedStr] = append(dict[sortedStr], val)
	}

	ans := make([][]string, 0, len(dict))
	for _, val := range dict {
		ans = append(ans, val)
	}

	return ans
}

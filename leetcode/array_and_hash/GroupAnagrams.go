package array_and_hash

import (
	"sort"
	"strconv"
)

func GroupAnagrams(strs *[]string) [][]string {
	dict := make(map[string][]string, len(*strs))

	for _, val := range *strs {
		freq := make([]int, 26)
		for _, ch := range val {
			freq[ch-'a']++
		}

		// Convert the frequency array to a string to use as a key
		key := ""
		for i, count := range freq {
			// Append the character and its count to the key
			key += string('a'+i) + strconv.Itoa(count)
		}

		// Append the string to the corresponding group in the map
		dict[key] = append(dict[key], val)
	}

	// Prepare the final result by converting the map values to a slice
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

// optimised
func GroupAnagram3(strs *[]string) [][]string {
	freqs := make(map[[26]byte][]string, len(*strs))

	for _, s := range *strs {
		g := [26]byte{}

		for _, c := range s {
			g[c-'a']++
		}

		freqs[g] = append(freqs[g], s)
	}

	result := [][]string{}

	for _, v := range freqs {
		result = append(result, v)
	}

	return result
}

func GroupAnagram4(strs *[]string) [][]string {
	freqs := make(map[[26]byte][]string, len(*strs))

	for _, arr_values := range *strs {
		key := [26]byte{}

		for _, val := range arr_values {
			key[val-'a']++
		}

		freqs[key] = append(freqs[key], arr_values)
	}

	result := make([][]string, 0, len(freqs))

	for _, val := range freqs {
		result = append(result, val)
	}

	return result
}

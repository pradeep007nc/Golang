package array_and_hash

import (
	"fmt"
	"sort"
)

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
func TopKFreqentElements(nums []int, k int) []int {
	dict := make(map[int]int)

	for _, val := range nums {
		dict[val]++
	}

	tempArr := make([][]int, len(nums)+1)

	for key, val := range dict {
		tempArr[val] = append(tempArr[val], key)
	}

	fmt.Println(tempArr)

	// index := 0
	ans := make([]int, 0, len(nums))
	for i := len(tempArr) - 1; i >= 0; i-- {
		if len(ans) == k {
			return ans
		}
		j := 0
		for j < len(tempArr[i]) && len(ans) < k {
			ans = append(ans, tempArr[i][j])
			j++
		}
	}

	return ans
}

func TopKFrequent2(nums []int, k int) []int {
	dict := make(map[int]int, len(nums))

	for _, val := range nums {
		dict[val]++
	}

	// Create a list to store frequency and corresponding values
	freqList := make([][]int, len(nums)+1)
	for key, val := range dict {
		freqList[val] = append(freqList[val], key)
	}

	ans := make([]int, 0)
	// Iterate in reverse order to get the elements with higher frequency first
	for i := len(freqList) - 1; i >= 0 && len(ans) < k; i-- {
		if freqList[i] != nil {
			ans = append(ans, freqList[i]...)
		}
	}

	return ans[:k]
}

func TopKFreqentElements3(nums []int, k int) []int {
	// Create a map to store the frequency of each element in the nums slice.
	dict := make(map[int]int, len(nums))

	// Count the frequency of each element in the nums slice.
	for _, val := range nums {
		dict[val]++
	}

	// Create a temporary array to store elements based on their frequencies.
	tempArr := make([]int, 0, len(dict))

	// Fill the temporary array with elements based on their frequencies.
	for key := range dict {
		tempArr = append(tempArr, key)
	}

	// Sort tempArr based on frequencies in descending order.
	sort.Slice(tempArr, func(i, j int) bool {
		return dict[tempArr[i]] > dict[tempArr[j]]
	})

	// Initialize variables for the result.
	ans := make([]int, k)

	// Copy the top k elements from tempArr to ans.
	copy(ans, tempArr[:k])

	return ans
}

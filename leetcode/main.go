package main

import (
	"fmt"

	"dev.pradeep/packages/array_and_hash"
)

func main() {

	// problem 1
	// arr := []int{1, 2, 3, 4}
	// fmt.Print(array_and_hash.ContainsDuplicate(&arr))
	// fmt.Print(array_and_hash.ContainsDuplicate2(arr))

	// problem 2
	// valid anagram
	// str1 := "rat"
	// str2 := "tar"
	// fmt.Println(array_and_hash.ValidAnagram(str1, str2))

	// problem 3
	// group anagrams
	// data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// fmt.Println(array_and_hash.GroupAnagram4(&data))

	// problem 4
	//Top k freqent elemnets
	// nums := []int{1, 1, 2, 2, 3, 4}
	// k := 2
	// fmt.Println(array_and_hash.TopKFreqentElements(nums, k))

	// problem 5
	//product of array self
	nums := []int{1, 2, 3, 4}
	fmt.Println(array_and_hash.ProductExceptSelf(nums))
}

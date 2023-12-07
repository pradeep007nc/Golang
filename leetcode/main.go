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
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(array_and_hash.GroupAnagrams(&data))
}

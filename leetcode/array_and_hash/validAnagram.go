package array_and_hash

// beats 64 percent
// using map
func ValidAnagram(string1 string, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}

	dict1 := make(map[rune]int, len(string1))

	for i := range string1 {
		dict1[rune(string1[i])]++
	}

	for i := range string2 {
		dict1[rune(string2[i])]--
		if dict1[rune(string2[i])] < 0 {
			return false
		}
	}

	return true
}

// using array
func ValidAnagram2(string1 string, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}

	dict1 := make([]int, 26)

	for i := range string1 {
		dict1[string1[i]-'a']++
		dict1[string2[i]-'a']--
	}

	for _, val := range dict1 {
		if val != 0 {
			return false
		}
	}

	return true
}

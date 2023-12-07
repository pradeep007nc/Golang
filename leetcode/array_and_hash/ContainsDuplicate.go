package array_and_hash

// approach 1
type void struct{}

var member void

func ContainsDuplicate(arr *[]int) bool {
	set := make(map[int]void)
	for _, val := range *arr {
		if _, ok := set[val]; ok {
			return true
		}
		set[val] = member
	}
	return false
}

// approach 2
func ContainsDuplicate2(arr []int) bool {
	set := make(map[int]bool, len(arr))
	for _, val := range arr {
		if set[val] {
			return true
		}
		set[val] = true
	}
	return false
}

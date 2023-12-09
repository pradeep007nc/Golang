package array_and_hash

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		prod := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				prod *= nums[j]
			}
		}
		ans[i] = prod
	}

	return ans
}

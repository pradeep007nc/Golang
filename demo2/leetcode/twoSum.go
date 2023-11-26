package twoSum

import "math/rand"

func TwoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	dict := make(map[int]int)

	for index, val := range nums {
		compliment := target - val
		if _, ok := dict[compliment]; ok {
			ans[0] = dict[compliment]
			ans[1] = index
			return ans
		}
		dict[val] = index
	}

	return nil
}

func GenerateRandom(arrSize int, limit int) []int {
	arr := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		arr[i] = rand.Intn(limit + 1)
	}

	return arr
}

func TwoSumTime(times int, arrSize int) {
	for i := 0; i < times; i++ {
		randomArr := GenerateRandom(arrSize, i)
		target := rand.Intn(i + 1)

		_ = TwoSum(randomArr, target)
	}
}

package data_structures_algorithms

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result, length := make([][]int, 0), len(nums)
	sort.Ints(nums)
	var l, r, sum int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		l, r = i+1, length-1
		for l < r {
			sum = nums[l] + nums[r] + v
			if sum == 0 {
				result = append(result, []int{nums[l], nums[r], v})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}

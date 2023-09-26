package data_structures_algorithms

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0], result[length-1] = 1, 1

	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	postProduct := 1
	for i := length - 2; i >= 0; i-- {
		postProduct *= nums[i+1]
		result[i] *= postProduct
	}

	return result
}

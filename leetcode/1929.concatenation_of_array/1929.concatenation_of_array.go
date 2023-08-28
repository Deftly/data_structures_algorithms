package data_structures_algorithms

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func getConcatenationGeneric(nums []int) []int {
	const n = 2
	ans := make([]int, 0, len(nums)*n)
	for i := 0; i < n; i++ {
		ans = append(ans, nums...)
	}
	return ans
}

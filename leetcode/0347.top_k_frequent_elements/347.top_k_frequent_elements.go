package data_structures_algorithms

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	arr := make([][]int, len(nums)+1)

	for k, v := range m {
		arr[v] = append(arr[v], k)
	}

	result := make([]int, 0, k)

	for i := len(arr) - 1; i >= 0; i-- {
		for _, v := range arr[i] {
			if k > 0 {
				result = append(result, v)
				k--
			}
		}
	}
	return result
}

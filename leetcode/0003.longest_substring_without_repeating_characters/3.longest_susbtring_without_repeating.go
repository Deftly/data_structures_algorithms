package data_structures_algorithms

func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		if index, ok := indexes[s[right]]; ok && index >= left {
			left = index + 1
		}
		indexes[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

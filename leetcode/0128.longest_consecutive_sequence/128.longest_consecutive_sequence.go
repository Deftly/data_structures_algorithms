package data_structures_algorithms

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for _, n := range nums {
		numSet[n] = true
	}

	maxLen := 0

	for n := range numSet {
		if !numSet[n-1] {
			seqLen := 1
			for {
				if numSet[n+seqLen] {
					seqLen++
					continue
				}
				maxLen = max(seqLen, maxLen)
				break
			}
		}
	}

	return maxLen
}

func longestConsecutiveV2(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	maxLen := 0

	for n := range numSet {
		if _, ok := numSet[n-1]; !ok {
			seqLen := 1
			for {
				if _, ok := numSet[n+seqLen]; ok {
					seqLen++
					continue
				}
				maxLen = max(seqLen, maxLen)
				break
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

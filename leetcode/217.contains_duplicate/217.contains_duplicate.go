package containsduplicate

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	record := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
	}
	return false
}

package data_structures_algorithms

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alphabet := make([]int, 26)
	for _, v := range s {
		alphabet[v-'a']++
	}
	for _, v := range t {
		alphabet[v-'a']--
	}
	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramRune(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	isAnagram := map[rune]int{}
	for _, v := range s {
		isAnagram[v]++
	}
	for _, v := range t {
		isAnagram[v]--
	}
	for _, value := range isAnagram {
		if value != 0 {
			return false
		}
	}
	return true
}

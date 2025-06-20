package validanagram

func ValidAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	tracker := [26]int{}
	//tracker := make([]int, 26) use make to create slices

	for _, v := range s {
		tracker[v-'a']++
	}

	for _, v := range t {
		tracker[v-'a']--
	}

	for _, v := range tracker {
		if v != 0 {
			return false
		}

	}

	return true
}

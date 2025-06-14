package groupanagrams

func GroupAnagrams(strs []string) [][]string {
	track := make(map[[26]int][]string)

	for _, str := range strs {

		key := [26]int{}

		for _, char := range str {

			key[char-'a']++

		}

		track[key] = append(track[key], str)
	}

	results := [][]string{}

	for _, v := range track {
		results = append(results, v)
	}

	return results

}

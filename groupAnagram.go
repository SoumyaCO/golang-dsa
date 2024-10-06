package main

/*
strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
*/

func groupAnagram(strs []string) [][]string {
	sc := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int
		for _, letter := range word {
			count[letter-'a']++
		}

		sc[count] = append(sc[count], word)
	}

	var result [][]string
	for _, value := range sc {
		result = append(result, value)
	}
	return result
}

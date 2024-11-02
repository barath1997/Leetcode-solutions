// using sorting , found from solutions
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)

		if group, found := anagramMap[sortedWord]; found {
			anagramMap[sortedWord] = append(group, word)
		} else {
			anagramMap[sortedWord] = []string{word}
		}
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}


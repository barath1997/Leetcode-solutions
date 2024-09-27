func commonChars(words []string) []string {

	n := len(words)
    result := make([]string, 0)
    if n == 1 {
      for _,v := range words[0] {
        result = append(result,string(v))
      }
      return result
    }
	
	lettersArr := make([]map[string]int, n)
	for i := 0; i < n; i++ {
		lettersArr[i] = make(map[string]int, 0)
	}
	for i := 0; i < n; i++ {
		newMap := make(map[string]int)
		for j := 0; j < len(words[i]); j++ {
			newMap[string(words[i][j])]++
		}
		lettersArr[i] = newMap
	}

	minCountMap := make(map[string]int, 0)
	firstWordMap := lettersArr[0]
    deletedLetters := make(map[string]bool,0) 

	for i := 1; i < n; i++ {
		for letters, freqency := range firstWordMap {
			if freq := lettersArr[i][letters]; freq >= 1 {
				if minCountMap[letters] != 0 {
					minCountMap[letters] = min(freq, freqency, minCountMap[letters])
				} else if deletedLetters[letters] != true {
					minCountMap[letters] = min(freq, freqency)
				}
			} else {
                deletedLetters[letters] = true
				minCountMap[letters] = 0
			}
		}
	}
	for alphabet, freq := range minCountMap {
		for freq > 0 {
			result = append(result, alphabet)
			freq--
		}
	}
	return result
}
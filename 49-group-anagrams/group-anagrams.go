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

// TC : O(N2 * K)
/*func groupAnagrams(strs []string) [][]string {

	n := len(strs)
	if n == 1 {
	   return [][]string{[]string{strs[0]}}
	}

	freqMap := make(map[string]int,0)
	result := make([][]string,0)
    hashedIndexes := make([]int,0)

	for i:=0;i<n;i++ {
		if contains(hashedIndexes,i) {
			continue
			} else if i == n-1 {
			result = append(result, []string{strs[i]});continue
		}
		deleteKeys(freqMap)
		for j:=0;j<len(strs[i]);j++ {
		   freqMap[string(strs[i][j])]++
		}

		copyMap := make(map[string]int,0)
		copyMaps(copyMap,freqMap)
		temp := make([]string,0)
		temp = append(temp, strs[i])
		hashedIndexes = append(hashedIndexes, i)
		for b:=i+1;b<n;b++ {
		   for k:=0;k<len(strs[b]);k++ {
			   copyMap[string(strs[b][k])]--
		   }
		   if allZeros(copyMap) {
              temp = append(temp, strs[b])
			  hashedIndexes = append(hashedIndexes, b)
		   }
		   copyMaps(copyMap,freqMap)
		}
		result = append(result, temp)
		temp = []string{}

	}
	return result

}

func contains(arr []int, idx int) bool{

    for _,v := range arr {
		if v == idx {
			return true
		}
	}
    return false
}

func copyMaps(dst , src map[string]int) {
    for key,val := range src {
        dst[key] = val
    }

	for key,_ := range dst {
		if _,ok := src[key];!ok {
            delete(dst,key)
		}
	}
}

func deleteKeys(m map[string]int) {
	for key,_ := range m {
		delete(m,key)
	}
}

func allZeros(m map[string]int) bool {

 for _,v := range m {
   if v != 0 {
	   return false
   }
 }
 return true
}*/
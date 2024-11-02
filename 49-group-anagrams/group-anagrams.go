//
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
	   return [][]string{[]string{strs[0]}}
	}

    anagramGroupsMap := make(map[string][]string, 0)
    result := make([][]string,0)
    for _,v := range strs {
        b := []byte(v)
        slices.Sort(b)
        sortedStr := string(b)

        anagram := anagramGroupsMap[sortedStr]
        anagram = append(anagram,v)
        anagramGroupsMap[sortedStr] = anagram
    }

    for _,groups := range anagramGroupsMap {
        result = append(result,groups)
    }
    
    return result
}

// using sorting , found from solutions, TC : O(n *k *log(k))
/*func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string, 0)
	result := make([][]string, 0)
	for _, v := range strs {
		s := strings.Split(v, "")
		sort.Strings(s)
		s1 := strings.Join(s, "")
		if _, ok := m[s1]; ok {
			m[s1] = append(m[s1], v)
		} else {
			m[s1] = []string{v}
		}
	}

	for _, val := range m {
		result = append(result, val)
	}

	return result
}*/

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
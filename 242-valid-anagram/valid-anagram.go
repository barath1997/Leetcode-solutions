func isAnagram(s string, t string) bool {

	stMap := make(map[rune]int, len(s))

	for _, v := range s {
		stMap[v]++
	}

	for _, v := range t {
		if val := stMap[v]; val > 0 {
			stMap[v]--
		} else {
			return false
		}
	}

    for _,v := range stMap {
        if v > 0 {
            return false
        }
    }

	return true

}
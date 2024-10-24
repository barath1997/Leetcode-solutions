func isAnagram(s string, t string) bool {
  
      if len(s) != len(t) {return false}
    if s == t {return true}

	stMap := make(map[string]int, len(s))

    for i:=0;i<len(s);i++ {
        stMap[string(s[i])]++
        stMap[string(t[i])]--
    }

    for _,v := range stMap {
        if v != 0 {
            return false
        }
    }

    return true

}


/*func isAnagram(s string, t string) bool {

    if len(s) != len(t) {return false}
    if s == t {return true}

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

	return true

    // second solution is based on sorting
    // we sort the string and compare them.

}*/
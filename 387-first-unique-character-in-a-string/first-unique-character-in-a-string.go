func firstUniqChar(s string) int {

    m := make(map[rune]int,len(s))

    for _,v := range s {
		m[v]++
	}

	for idx,v := range s {
		if m[v] == 1 {return idx}
	}
     
	return -1
}
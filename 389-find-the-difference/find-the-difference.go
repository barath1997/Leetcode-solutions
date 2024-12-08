// USING XOR OPERATOR
func findTheDifference(s string, t string) byte {
    
    var result rune
	for _,v := range s {
		result = result ^ v
	}

	for _,v := range t {
		result = result ^ v
	}

    return []byte(string(result))[0]

}

// TC : O(2N) , SC : O(N)
/*func findTheDifference(s string, t string) byte {

	if s == "" {
		return []byte(t)[0]
	}

	m := make(map[byte]int, len(t))

	for i := 0; i < len(t); i++ {
		if i < len(s) {
			m[s[i]]++
		}

		m[t[i]]--
	}

	for char, v := range m {
		if v < 0 {
			return char
		}
	}

	return []byte("")[0]
}*/
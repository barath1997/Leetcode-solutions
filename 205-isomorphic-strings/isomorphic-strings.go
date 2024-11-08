func isIsomorphic(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	m := make(map[string]string, 0)

    // we put first character of s and t strings in the map
	m[string(s[0])] = string(t[0])

    //we traverse from index 1(i.e next character in the string) "NOTE : len(s) == len(t) given in constrains"
	for i := 1; i < len(s); i++ {
		if val, ok := m[string(s[i])]; ok {

            // if key,value pair exists in map, then we have to check in the value == t[i] to check if key   
            // and values are mapped correctly.
			if val != string(t[i]) {
				return false
			}
		} else {
			count := 0
            // check for the value in the map
			for _, v := range m {
                // if the value already exists in the map then return false
				if v == string(t[i]) {
					return false
				}
                // counts keeps track of end of the map traversal
				count++
			}

            // count == len(m) indicates that the value does not exist in map so we add the key,value to it
			if count == len(m) {
				m[string(s[i])] = string(t[i])
				count = 0
			}
		}
	}

	return true

}
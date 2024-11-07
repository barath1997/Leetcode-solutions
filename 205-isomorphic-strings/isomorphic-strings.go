func isIsomorphic(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	m := make(map[string]string, 0)
	m[string(s[0])] = string(t[0])

	for i := 1; i < len(s); i++ {
		if val, ok := m[string(s[i])]; ok {
			if val != string(t[i]) {
				return false
			}
		} else {
			count := 0
			for _, v := range m {
				if v == string(t[i]) {
					return false
				}
				count++
			}
			if count == len(m) {
				m[string(s[i])] = string(t[i])
				count = 0
			}
		}
	}

	return true

}
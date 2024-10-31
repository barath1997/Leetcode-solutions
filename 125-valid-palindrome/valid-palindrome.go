func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i := 0
	j := len(s) - 1

	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
		} else if !isAlphanumeric(s[j]) {
			j--
		} else if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		} else {
			i++
			j--
		}
	}

	return true
}
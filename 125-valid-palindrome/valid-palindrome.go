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
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}

		for i < j && !isAlphanumeric(s[j]) {
			j--
		}

		if i >= j {
			return true
		}

		if i < j && strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
func isPalindrome(s string) bool {
    for _,c := range s {
		if val := c ; !check(int(val)) {
          s = strings.ReplaceAll(s,string(c),"")
		} 
	}

	s = strings.ToLower(s)
	
	return validPalindrome(s)
}

func check(char int) bool {
	if char >= 97 && char <= 122 {
		return true
	} else if char >= 65 && char <= 90 {
		return true
	} else if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func validPalindrome(s string) bool {

     begin := 0
	 end := len(s)-1

	 for begin <= end {
		if s[begin] != s[end] {
			return false
		}
		begin++;end--
	 }

	 return true
}
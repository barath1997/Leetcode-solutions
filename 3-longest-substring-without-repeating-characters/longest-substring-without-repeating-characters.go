func lengthOfLongestSubstring(s string) int {

    if s == "" {return 0}
    if len(s) == 1 {return 1}

    existing := make(map[string]bool,0)
    maxLen := math.MinInt
    subString := ""

    for i:=0;i<len(s)-1;i++ {
		if existing[string(s[i])] {
            continue
		} else {
			subString += string(s[i])
			existing[string(s[i])] = true
		}

        for j:=i+1;j<len(s);j++ {
            if !existing[string(s[j])] {
                subString += string(s[j])
                existing[string(s[j])] = true
            } else {
				maxLen = max(maxLen,len(subString))
                existing = make(map[string]bool,0)
                subString = ""
                break
            }
        }
    }
    
    return max(maxLen,len(subString))
}
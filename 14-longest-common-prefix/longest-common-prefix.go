func longestCommonPrefix(strs []string) string {

	n := len(strs)
	if n == 1 {
		return strs[0]
	}
    str := ""
	first := strs[0]
	for idx, v := range first {
        i := 1
		for  ;i < n; i++ {
            if len(strs[i])-1 >= idx {
	    	if string(strs[i][idx]) != string(v) {
                 break
            }
        } else {
            break
        }
		}
        if i == n {
            str += string(v)
        } else {
            break
        }
	}
	return str
}
// https://www.youtube.com/watch?v=5X-c0ZzL9CI
func generateParenthesis(n int) []string {

	result := make([]string, 0)
	newCombo(0, 0, n, "", &result)
	return result
}

func newCombo(open, closed, total int, current string, result *[]string) {

	// base condition
	if len(current) == 2*total {
		*result = append(*result, current)
		return 
	}

	if open < total {
		newCombo(open+1, closed, total, current + "(", result)
	}

	if closed < open {
		newCombo(open, closed+1, total, current + ")", result)
	}

}


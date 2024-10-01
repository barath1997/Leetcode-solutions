// reference : https://www.youtube.com/watch?v=cgPqbpcIgoc
func multiply(num1 string, num2 string) string {

	s1, s2 := reverseString(num1), reverseString(num2)
	result := make([]int, len(s1)+len(s2))

	for i := 0; i < len(s1); i++ {
		var carry int = 0
		start := i
		for j := 0; j < len(s2); j++ {

			n1 := s1[i] - '0'
			n2 := s2[j] - '0'

			sum := (int(n1) * int(n2)) + carry + result[start]
			base := sum % 10
			carry = sum / 10

			result[start] = base
			start++
		}
		if carry != 0 {
			result[start] = carry
		}
	}

	i := len(result) - 1
	for i > 0 && result[i] == 0 {
		i--
	}

	ans := ""
	for i >= 0 {
		ans += fmt.Sprint(result[i])
		i--
	}

	return ans
}

func reverseString(s string) string {
	r := []rune(s)
	begin, end := 0, len(s)-1

	for begin < end {
		temp := r[begin]
		r[begin] = r[end]
		r[end] = temp
		begin++
		end--
	}

	return string(r)
}
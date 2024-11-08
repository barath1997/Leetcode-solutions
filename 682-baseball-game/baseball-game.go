type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(num int) {
	*s = append(*s, num) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return math.MinInt, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (int, bool) {
	if !s.IsEmpty() {
		return (*s)[len(*s)-1], true
	}
	return math.MinInt, false
}

func calPoints(operations []string) int {

	var s Stack
	sum := 0

	for _, v := range operations {

		switch string(v) {
		case "+":
			val, _ := s.Pop()
			val2, _ := s.Peek()
			sum := val + val2
			s.Push(val)
			s.Push(sum)
		case "C":
			s.Pop()
		case "D":
			val, _ := s.Peek()
			s.Push(val * 2)
		default:
			num, _ := strconv.Atoi(string(v))
			s.Push(num)
		}
	}

	for _, v := range s {
		sum += v
	}
	return sum

}
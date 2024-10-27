type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (string, bool) {
	if !s.IsEmpty() {
		return (*s)[len(*s)-1], true
	}
	return "", false
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var stack Stack
	pushCount, popCount := 0, 0
	for _, c := range s {
		v := string(c)
		if v == "(" || v == "{" || v == "[" {
			stack.Push(v)
			pushCount++
		} else {
			if v == ")" {
				if val, ok := stack.Peek(); ok {
					if val == "(" {
						stack.Pop()
					} else {
						return false
					}
				} else {
					return false
				}
			} else if v == "}" {
				if val, ok := stack.Peek(); ok {
					if val == "{" {
						stack.Pop()
					} else {
						return false
					}
				} else {
					return false
				}
			} else if v == "]" {
				if val, ok := stack.Peek(); ok {
					if val == "[" {
						stack.Pop()
					} else {
						return false
					}
				} else {
					return false
				}
			}
			popCount++
		}
	}

	return pushCount == popCount
}
func dailyTemperatures(temperatures []int) []int {

  n := len(temperatures)
  if n == 1 {return []int{0}}
  var stack Stack
  result := make([]int,n)

  for i:=0;i<n;i++ {
    result[i] = 0
  }

  for i:=n-1;i>=0;i-- {

     for !stack.IsEmpty() && temperatures[stack.Peek()] <= temperatures[i] {
        stack.Pop()
     }

     if !stack.IsEmpty() {
        result[i] = stack.Peek() - i
     } 

     stack.Push(i)
  }

  return result 
}

// func dailyTemperatures(temperatures []int) []int {
//     n := len(temperatures)
//     var stack Stack
//     result := make([]int,n)
//     for i:=n-1;i>=0;i-- {
// 		idx := stack.Peek()
//         for !stack.IsEmpty() && temperatures[idx] <= temperatures[i] {
//             stack.Pop()
//             idx = stack.Peek()
//         }
//         if stack.IsEmpty() {
//             result[i] = 0
//         } else {
//             result[i] = stack.Peek() - i
//         }
//         stack.Push(i)
//     }

//     return result
// }

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(num int) {
	*s = append(*s, num) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return math.MinInt
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

func (s *Stack) Peek() int {
	if !s.IsEmpty() {
		return (*s)[len(*s)-1]
	}
	return math.MinInt
}
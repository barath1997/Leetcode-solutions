// without using stack
func minAddToMakeValid(s string) int {
    open,close := 0,0

    for _,v := range s {
        val := string(v)

        if val == "(" {
            close++
        } else if close > 0 {
            close--
        } else {
            open++
        }
    } 
    
    return open + close
}

/*func minAddToMakeValid(s string) int {
    
	openCount,closeCount := 0,0
	var stack Stack
 
	for _,v := range s {
	   val := string(v)
	   if val == "(" {
		 stack.Push(val)
		 closeCount++
	   } else {
		 bracket,_ := stack.Peek()
		 if bracket == "(" {
			 stack.Pop()
			 closeCount--
		 } else {
			 openCount ++
			 stack.Push(val)
		 }
	   }
 
	}
	return openCount + closeCount
 
 }
 
 type Stack []string
 
 func (s *Stack) IsEmpty() bool {
	 return len(*s) == 0
 }
 
 // Push a new value onto the stack
 func (s *Stack) Push(num string) {
	 *s = append(*s, num) // Simply append the new value to the end of the stack
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
 }*/
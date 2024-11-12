// STACK SOLUTION DISCUSSED IN CLASS
func nextGreaterElements(nums []int) []int {
   
   var stack Stack1
   n := len(nums)
   result := make([]int,n)

   for idx,_ := range nums {
      result[idx] = -1
   }

   for i:=0;i<2*n;i++ {
       
       for !stack.IsEmpty() && nums[stack.Peek()] < nums[i%n] {
         result[stack.Peek()] = nums[i%n]
         stack.Pop()
       }
       
       if i<n {
        stack.Push(i)
       }
   }

   return result
}

type Stack1 []int

func (s *Stack1) IsEmpty() bool{
    return len(*s) == 0
}

func (s *Stack1) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack1) Peek() int {
	if len(*s) != 0 {
		dd := *s
		return dd[len(*s)-1]
	}
	return math.MinInt
}

func (s *Stack1) Pop() {
	if len(*s) != 0 {
		dd := *s
		*s = dd[:len(*s)-1]
	}
}

// MONOTONIC STACK
// func nextGreaterElements(nums []int) []int {
// 	initialLength := len(nums)
// 	nums = append(nums, nums...)

// 	ans, stack := make([]int, initialLength), make([]int, 0)
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
// 			stack = stack[:len(stack)-1]
// 		}
// 		if i < initialLength {
// 			if len(stack) != 0 {
// 				ans[i] = stack[len(stack)-1]
// 			} else {
// 				ans[i] = -1
// 			}
// 		}
// 		stack = append(stack, nums[i])
// 	}
// 	return ans
// }

// BRUTE FORCE
// func nextGreaterElements(nums []int) []int {
// 	length := len(nums)
// 	ans := make([]int, 0)
// 	for i := 0; i < length; i++ {
// 		for j := i; j < (i + length); j++ {
// 			if nums[j%length] > nums[i] {
// 				ans = append(ans, nums[j%length])
// 				break
// 			}
// 		}
// 		if len(ans) == i {
// 			ans = append(ans, -1)
// 		}
// 	}
// 	return ans
// }
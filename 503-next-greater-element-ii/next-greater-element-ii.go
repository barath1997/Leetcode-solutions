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
func nextGreaterElements(nums []int) []int {
	length := len(nums)
	ans := make([]int, 0)
	for i := 0; i < length; i++ {
		for j := i; j < (i + length); j++ {
			if nums[j%length] > nums[i] {
				ans = append(ans, nums[j%length])
				break
			}
		}
		if len(ans) == i {
			ans = append(ans, -1)
		}
	}
	return ans
}
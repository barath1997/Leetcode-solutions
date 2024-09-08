func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {return nums[0]}
    
    // moore's best voting algorithm
   counts := make(map[int]int,n)

   for _,val := range nums {
    counts[val]++
   }

   for val,c := range counts {
    if c > n/2 {
        return val
    }
   }
    
    return -1

}
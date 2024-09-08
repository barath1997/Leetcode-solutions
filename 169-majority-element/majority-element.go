func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {return nums[0]}

   counts := make(map[int]int,n)

   start,end := 0,n-1

   for start <= end {
    counts[nums[start]]++
    counts[nums[end]]++
    
    if c,_ := counts[nums[start]];c > n/2 {
      return nums[start]
    } 
    if c,_ := counts[nums[end]];c > n/2 {
        return nums[end]
    } 
    start++;end--
   }
 
   return -1

}
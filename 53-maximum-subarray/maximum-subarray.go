func maxSubArray(nums []int) int {
    
   n := len(nums)
   maxSum,currSum := math.MinInt,0

   for i:=0;i<n;i++ {
    
     currSum += nums[i]
     maxSum = max(maxSum,currSum)

     if currSum < 0 { //reset currSum
        currSum = 0
     }
   }

   return maxSum

}
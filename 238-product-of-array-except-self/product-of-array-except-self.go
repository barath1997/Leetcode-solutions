func productExceptSelf(nums []int) []int {
    
    n := len(nums)
   prefixProductArr := make([]int,n)
   result := make([]int,n)
   prefixProductArr[0] = nums[0] 
   for i:=1;i<n;i++ {
    prefixProductArr[i] = nums[i] * prefixProductArr[i-1]
   }

   reversePrefixProductArr := make([]int,n)
   reversePrefixProductArr[n-1] = nums[n-1]

   for i:=n-2;i>0;i-- {
      reversePrefixProductArr[i] = reversePrefixProductArr[i+1] * nums[i]
   }

   for i:=0;i<n;i++ {
     if i == 0 {
        result[i] = reversePrefixProductArr[i+1]
     } else if i == n-1{
        result[i] = prefixProductArr[i-1]
     } else {
        result[i] = prefixProductArr[i-1] * reversePrefixProductArr[i+1]
     }

   }

       return result

}
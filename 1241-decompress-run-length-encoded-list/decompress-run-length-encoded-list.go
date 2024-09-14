func decompressRLElist(nums []int) []int {

    n := len(nums)
    result := make([]int,0)
    for i:=0;i<n-1;i+=2 {
      newArr := make([]int,nums[i])
      for j:=0;j<nums[i];j++ {
        newArr[j] = nums[i+1]
      }
      result = append(result, newArr...)
    }
     
     return result
}
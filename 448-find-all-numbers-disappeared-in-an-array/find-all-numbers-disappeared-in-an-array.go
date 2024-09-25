// TC : O(n) , SC : O(n)
/*func findDisappearedNumbers(nums []int) []int {

    n := len(nums)
    result := make([]int,0)
    existsMap := make(map[int]bool,n)

    for _,v := range nums {
        existsMap[v] = true
    }

    for i:=1;i<=n;i++ {
        if existsMap[i] != true {
           result = append(result,i)
        }
    }
    
    return result
}*/

// Optimised TC : O(n) , SC : O(1) no extra space

func findDisappearedNumbers(nums []int) []int {

   result := make([]int,0)
   n := len(nums) 
   for i:=0;i<n;i++ {
     for nums[nums[i]-1] != nums[i] {
        nums[nums[i]-1],nums[i] = nums[i],nums[nums[i]-1]
     }
   }

   for i:=1;i<=n;i++ {
     if i != nums[i-1] {
        result = append(result,i)
     }
   }
   
   return result
}
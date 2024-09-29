func singleNumber(nums []int) int {
   
    n := len(nums) 
    sort.Ints(nums)
    for i:=1;i<n;i+=3 {
        if nums[i] == nums[i-1] && nums[i] == nums[i+1] {
            continue
        } else {
            return nums[i-1]
        }

    }

    return nums[n-1]
}
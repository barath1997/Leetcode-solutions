func moveZeroes(nums []int)  {
    
    n := len(nums)

    f := 0
    s := 1
    
    // since we have to maintain relative order, we dont use binary search technique
    for s < n {
        if nums[f] == 0 && nums[s] != 0 {
            temp := nums[f]
            nums[f] = nums[s]
            nums[s] = temp
        } else if nums[f] !=0 && nums[s] !=0 {
            f++;s++
        } else if nums[f] == 0 && nums[s] == 0 {
            s++
        } else {
            f++
        }
    }  
}
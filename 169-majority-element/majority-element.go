func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {return nums[0]}
    
    // moore's best voting algorithm
    majorityElem := nums[0]
    count := 1
    
    for i:=1;i<n;i++ {
        if count == 0 {majorityElem = nums[i]}
        if nums[i] == majorityElem {
            count++
        } else {count--}

        
    }
    
    return majorityElem

}
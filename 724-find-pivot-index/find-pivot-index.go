func pivotIndex(nums []int) int {

    eqIndex := 0
    n := len(nums)

    if n == 0 {return -1}
    if n == 1 {return 0}

    for eqIndex < n {
        rSum,lSum := 0,0
        for i:=n-1;i>eqIndex;i-- {
            rSum += nums[i]
        }

        for i:=0;i<eqIndex;i++ {
            lSum += nums[i]
        }

        if lSum == rSum {return eqIndex}
        eqIndex++
    }

    return -1
    
}
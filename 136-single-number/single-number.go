func singleNumber(nums []int) int {
     // optimised solution
    n := len(nums)
    res := nums[0]
    for i:=1;i<n;i++ {
        res ^= nums[i]
    }
    
    return res

}

 // brute force solution
    /*m := make(map[int]int,n)

    for _,v := range nums {
        m[v]++
    }
    
    for num,count := range m {
        if count == 1 {
            return num
        }
    }
    return 0*/
func singleNumber(nums []int) int {
    
    
    n := len(nums)
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

    // optimised solution
    res := nums[0]
    for i:=1;i<n;i++ {
        res ^= nums[i]
    }
    
    return res

}
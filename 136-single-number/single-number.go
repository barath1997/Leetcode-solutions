func singleNumber(nums []int) int {
    
    // brute force solution
    n := len(nums)
    m := make(map[int]int,n)

    for _,v := range nums {
        m[v]++
    }
    
    for num,count := range m {
        if count == 1 {
            return num
        }
    }
    return 0
}
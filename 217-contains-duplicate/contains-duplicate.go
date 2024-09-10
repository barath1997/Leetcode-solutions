func containsDuplicate(nums []int) bool {
    
    counterMap := make(map[int]int,len(nums))

    for _,val := range nums {
        counterMap[val]++
        if count := counterMap[val]; count > 1 {
            return true
        }
    }
    return false
}
func singleNumber(nums []int) []int {

    // map based soln : TC : O(n) , SC : O(n)
    counts := make(map[int]int,len(nums))
    result := make([]int,2)
    for _,v := range nums {
        counts[v]++
    }
    
    idx := 0
    for val,count := range counts {
        if count == 1 {
          result[idx] = val
          idx++
        }
    }

    return result
}
func majorityElement(nums []int) []int {
    
    n := len(nums)
    counts := make(map[int]int,n)
    result := make([]int,0)

    for _,value := range nums {
        counts[value]++
    }

    for key,value := range counts {
        if value > n/3 {
            result = append(result,key)
        }
    }

    return result

}
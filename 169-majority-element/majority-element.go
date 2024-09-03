func majorityElement(nums []int) int {
    n := len(nums)
    counts := make(map[int]int,n)

    for _,v := range nums {
        counts[v]++
    }

    for key,value := range counts {
        if value > n/2 { 
            return key
        }
    }

    return 0
}
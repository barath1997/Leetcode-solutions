func intersection(nums1 []int, nums2 []int) []int {

    m := make(map[int]int,min(len(nums1),len(nums2)))
    result := make([]int,0)
    for _,val := range nums1 { 
        m[val]++
    }

    for _,val := range nums2 {
        if _,ok := m[val];ok {
           result = append(result,val)
           delete(m,val)
        }
    }
    
    return result
}
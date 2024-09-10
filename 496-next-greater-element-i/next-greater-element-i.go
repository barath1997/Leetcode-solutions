func nextGreaterElement(nums1 []int, nums2 []int) []int {
     result := make([]int,len(nums1))
     nums2Map := make(map[int]int,len(nums2))
        for idx,val := range nums2 {
           nums2Map[val] = idx
        }
        
        for j,val := range nums1 {
            if idx , ok := nums2Map[val];ok {
                if idx + 1 == len(nums2) {result[j] = -1; continue}
                for i:=idx+1;i<len(nums2);i++ {
                    if nums2[i] > val { result[j] = nums2[i]; break } else {result[j] = -1}
                }
            }
        }

        return result
}
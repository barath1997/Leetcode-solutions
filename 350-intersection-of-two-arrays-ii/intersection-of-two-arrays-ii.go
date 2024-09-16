func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2))
	result := make([]int, 0)
	for _, val := range nums2 {
		m[val]++
	}
    
	for _, val := range nums1 {
		if idx, ok := m[val]; ok {
            if idx >= 1 {
			result = append(result, val)
			m[val]--
		   }
	}
    }

	return result
}
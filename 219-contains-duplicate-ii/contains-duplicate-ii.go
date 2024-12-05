func containsNearbyDuplicate(nums []int, k int) bool {

	n := len(nums)
    maps := make(map[int]int,n)

    for idx,val := range nums {
        if v,ok := maps[val];ok {
            if math.Abs(float64(v-idx)) <= float64(k) {
                return true
            }
            maps[val] = idx
        }
        maps[val] = idx
    }
	
	return false
}
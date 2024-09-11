func summaryRanges(nums []int) []string {
    
    n := len(nums)
    if n == 0 {return []string{}}
    if n == 1 {return []string{strconv.Itoa(nums[0])}}

	first, count := nums[0], 0
   
	result := make([]string, 0)

	for i := 1; i < n; i++ {
		if nums[i] == first+1 {
			count++
		} else if count != 0 {
			result = append(result, strconv.Itoa(nums[i-count-1])+"->"+strconv.Itoa(nums[i-1]))
			count = 0
		} else if count == 0 {
			result = append(result, strconv.Itoa(first))
		}

		if count != 0 && i == n-1 {
			result = append(result, strconv.Itoa(nums[i-count])+"->"+strconv.Itoa(nums[i]))
		}

		if count == 0 && i == n-1 {
			result = append(result, strconv.Itoa(nums[i]))
		}

		first = nums[i]
	}

	return result
}
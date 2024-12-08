func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	sort.Ints(nums)
	fmt.Printf("arr : %d\n",nums)
	maxLen := math.MinInt
	count := 0
	minElement := nums[0]

	duplicateNumbers := make(map[int]bool, 0)

	for i := 0; i < n; {
		check := duplicateNumbers[nums[i]]
		if check {
			i++
			continue
		} else {
			duplicateNumbers[nums[i]] = true
		}

		if nums[i] == minElement {
			count++
			i++
			minElement++
		} else {
			maxLen = max(maxLen, count)
			count = 1
			minElement = nums[i]+1
			i++
		}
		maxLen = max(maxLen, count)
	}

	return maxLen

}
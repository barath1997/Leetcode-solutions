func removeElement(nums []int, val int) int {

	count := 0

	if val > 50 {
		return len(nums)
	}

	for idx, v := range nums {
		if v != val {
			count++
		} else {
			nums[idx] = 101
		}
	}

	timesToShift := len(nums)-count
    for i:=0;i<len(nums);i++ {
		if timesToShift < 1 {
			break
		}
        if nums[i] == 101 {
            leftShift(nums,i);i--
			timesToShift--
        }
    }
    
    return count
}

func leftShift(arr []int,leftIdx int) {
    for i:=leftIdx;i<len(arr)-1;i++ {
        arr[i] = arr[i+1]
    }
}
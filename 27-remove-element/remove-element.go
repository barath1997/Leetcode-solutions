func removeElement(nums []int, val int) int {
 
 // using 2 pointers
 start,end := 0,len(nums)-1

 for start<=end {
    if nums[start] == val && nums[end] != val {
        temp := nums[start]
        nums[start] = nums[end]
        nums[end] = temp
        start++;end--
    } else if nums[start] != val && nums[end] == val {
        start++;end--
    } else if nums[start] == val && nums[end] == val {
        end--
    } else {
        start++
    }
 }

 count := 0

 for _,v := range nums {
    if v != val {
        count++
    }
 }

 return count

}

/*func removeElement(nums []int, val int) int {

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
}*/

func leftShift(arr []int,leftIdx int) {
    for i:=leftIdx;i<len(arr)-1;i++ {
        arr[i] = arr[i+1]
    }
}
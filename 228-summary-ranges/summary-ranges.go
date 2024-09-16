func summaryRanges(nums []int) []string {
    
    n := len(nums)
    if n == 0 {return []string{}}
    if n == 1 {return []string{strconv.Itoa(nums[0])}}
	result := make([]string, 0)

	for i := 0; i < n; i++ {
    
       j := i

       for j+1 < n && nums[j+1] == nums[j]+1 {
        j++
       }

       if j > i {
        result = append(result,strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j]))
       } else {
         result = append(result,strconv.Itoa(nums[i]))
       }
       
       // reset
       i = j
       
    }

	return result
}
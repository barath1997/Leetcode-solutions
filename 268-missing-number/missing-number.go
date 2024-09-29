func missingNumber(nums []int) int {
    // brute force
    sort.Ints(nums)
    noOfIterations := nums[len(nums)-1] - nums[0] + 1
    for i:=0;i<noOfIterations;i++ {
        if i != nums[i] {
            return i
        }
    }
    return noOfIterations
}
func missingNumber(nums []int) int {
    // brute force
    /*sort.Ints(nums)
    noOfIterations := nums[len(nums)-1] - nums[0] + 1
    for i:=0;i<noOfIterations;i++ {
        if i != nums[i] {
            return i
        }
    }
    return noOfIterations*/

    // optimised
    sum := 0
    for _,v := range nums {
        sum+=v
    }
    n := len(nums)
    //sum of numbers form 0...n
    expectedSum := ((n+1)*n)/2
    return expectedSum - sum
}
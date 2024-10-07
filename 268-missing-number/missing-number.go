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
    /*sum := 0
    for _,v := range nums {
        sum+=v
    }
    n := len(nums)
    //sum of numbers from 0...n
    expectedSum := ((n+1)*n)/2
    return expectedSum - sum*/

    // using bit maniluation
    /*sort.Ints(nums)
    noOfIterations := nums[len(nums)-1] - nums[0] + 1
    for i:=0;i<noOfIterations;i++ {
        if i ^ nums[i] != 0 {
            return i
        }
    }
    return noOfIterations
    */

    // more optimised using bit manipulation and xor
    result := 0 ; i := 0
    for i=0;i<len(nums);i++ {
        result = result ^ i ^ nums[i]
    }

    return result ^ i
}
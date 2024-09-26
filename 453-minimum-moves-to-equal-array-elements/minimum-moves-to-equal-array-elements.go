func minMoves(nums []int) int {
    n := int32(len(nums))
    var sum int32
    min := int32(nums[0])
    var i int32 = 0
    for i=0;i<n;i++ {
        sum+=int32(nums[i])
        if int32(nums[i]) < min {
            min = int32(nums[i])
        }
    }
    return int(sum - (n*min))
} 

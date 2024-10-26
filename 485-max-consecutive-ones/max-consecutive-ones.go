func findMaxConsecutiveOnes(nums []int) int {
    maxOnes := math.MinInt
    ones := 0
    for _,v := range nums {
        if v == 1 {
            ones++
            maxOnes = max(maxOnes,ones)
        } else {
            ones = 0
            maxOnes = max(maxOnes,ones)
        }

    }
    return maxOnes
}
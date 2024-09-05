func majorityElement(nums []int) []int {
    count1, count2, num1, num2 := 0, 0, 0, 0

    // We can have only 2 numbers with count greater than len(nums)/3
    // Find those 2 possible numbers first
    for _, num := range nums {
        if num == num1 {
            count1++
        } else if num == num2 {
            count2++
        } else if count1 == 0 {
            num1, count1 = num, 1
        } else if count2 == 0 {
            num2, count2 = num, 1
        } else {
            count1--
            count2--
        }
    }

    // Verify count of those 2 possible numbers again
    count1, count2 = 0, 0
    for _, num := range nums {
        if num == num1 {
            count1++
        } else if num == num2 {
            count2++
        }
    }

    // Add those numbers if they satisfy conditions
    res := []int{}
    if count1 > len(nums) / 3 {
        res = append(res, num1)
    }
    if count2 > len(nums) / 3 {
        res = append(res, num2)
    }

    return res
}
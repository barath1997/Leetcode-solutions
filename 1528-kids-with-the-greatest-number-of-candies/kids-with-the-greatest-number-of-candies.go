func kidsWithCandies(candies []int, extraCandies int) []bool {
    maximum := math.MinInt
    result := make([]bool,len(candies))

    for _,v := range candies {
        if v > maximum {
            maximum = v
        }
    }

    for idx,v := range candies {
        if v+extraCandies >= maximum {
            result[idx] = true
        } else {
            result[idx] = false
        }
    }

    return result
}
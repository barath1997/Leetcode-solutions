func kidsWithCandies(candies []int, extraCandies int) []bool {
    maximum := math.MinInt
    result := make([]bool,len(candies))

    for _,v := range candies {
        if v > maximum {
            maximum = v
        }
    }

    for idx,v := range candies {
            result[idx] = v+extraCandies >= maximum
    }

    return result
}
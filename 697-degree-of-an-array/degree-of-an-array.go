func findShortestSubArray(nums []int) int {

    // find max frequency
    n := len(nums)
    frequency := make(map[int]int,n)

    for i:=0;i<n;i++ {
        frequency[nums[i]]++
    }

    maxFreq := 0
    for _,freq := range frequency {
        if freq > maxFreq {
            maxFreq = freq
        }
    }
    
    degrees := make([]int,0)
    for val,freq := range frequency {
        if freq == maxFreq {
            degrees = append(degrees,val)
        }
    }
    
    
    indexes := make([]int,0) // evenIndex:min , oddIndex:max
    for i:=0;i<len(degrees);i++ {
        minIndex,maxIndex := math.MaxInt,math.MinInt
        for j:=0;j<len(nums);j++ {
            if degrees[i] == nums[j] {
                minIndex = min(minIndex,j)
                maxIndex = max(maxIndex,j)
            }
        }
      indexes = append(indexes,minIndex,maxIndex)
    }

    result := math.MaxInt

    for i:=0;i<len(indexes)-1;i+=2 {
        result = min(result, indexes[i+1] - indexes[i])
    }

    return result + 1

}
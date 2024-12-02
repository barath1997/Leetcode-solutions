func largestNumber(nums []int) string {
    
    n := len(nums)
    arr := make([]string,n)

    for idx,v := range nums {
       arr[idx] = strconv.Itoa(v)
    }
    result := ""
    sort.SliceStable(arr,func(i,j int) bool {
        return arr[i] + arr[j] > arr[j] + arr[i]
    })

    if strings.HasPrefix(arr[0],"0")  {return "0"}

    for _,v := range arr {
        result += v
    } 

    return result
}
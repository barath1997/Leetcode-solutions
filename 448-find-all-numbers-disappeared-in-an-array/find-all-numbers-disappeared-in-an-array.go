func findDisappearedNumbers(nums []int) []int {

    n := len(nums)
    result := make([]int,0)
    existsMap := make(map[int]bool,n)

    for _,v := range nums {
        existsMap[v] = true
    }

    for i:=1;i<=n;i++ {
        if existsMap[i] != true {
           result = append(result,i)
        }
    }
    
    return result
}
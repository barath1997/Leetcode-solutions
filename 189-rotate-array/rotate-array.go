func rotate(nums []int, k int)  {


    n := len(nums)
    k = k%n
    begin,end := 0,n-1
    
    reverse(nums,begin,end)
    reverse(nums,begin,k-1)
    reverse(nums,k,end)
}

func reverse(arr []int,begin,end int) {
    for begin <= end {
        temp := arr[begin]
        arr[begin] = arr[end]
        arr[end] = temp
        begin++;end--
    }
}
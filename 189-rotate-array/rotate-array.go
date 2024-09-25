// optimised solution
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

// solution with O(n) extra space
/*func rotate(nums []int,k int) {
    
    n := len(nums)
    newArr := make([]int,n)

    for idx,val := range nums {
        newArr[idx] = val
    }

    for i:=0;i<n;i++ {
        nums[(i+k)%n] = newArr[i]
    }

}*/
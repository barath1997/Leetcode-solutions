// TC : O(N) , SC : O(1) , discussed in class
func sortColors(nums []int)  {
    
    end := len(nums)-1
    start := 0
    i := 0

    for i <= end {

        if nums[i] == 0 {       // move '0' to left
           swap(nums,i,start)
           i++
           start++
        } else if nums[i] == 1 {  // dont move '1' anywhere
            i++
        } else if nums[i] == 2 {  // move '2' to right
            swap(nums,i,end)
            end--
        }

    }
}

func swap(arr []int, p1,p2 int) {
    temp := arr[p1]
    arr[p1] = arr[p2]
    arr[p2] = temp
}
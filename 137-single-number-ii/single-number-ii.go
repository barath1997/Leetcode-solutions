func singleNumber(nums []int) int {
   
    /*n := len(nums) 
    sort.Ints(nums)
    for i:=1;i<n;i+=3 {
        if nums[i] == nums[i-1] && nums[i] == nums[i+1] {
            continue
        } else {
            return nums[i-1]
        }

    }

    return nums[n-1]*/

    //optimised solution : link : https://www.youtube.com/watch?v=5Bb2nqA40JY
    bucket1,bucket2 := 0,0
    for _,v := range nums {
       bucket1 = (bucket1 ^ v) & ^bucket2
       bucket2 = (bucket2 ^ v) & ^bucket1
    }

    return bucket1
}
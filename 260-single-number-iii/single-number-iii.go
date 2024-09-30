func singleNumber(nums []int) []int {
      
      // optimised solution
      xoredValue := 0 

      for _,v := range nums {
        xoredValue^=v
      }
      
      var val int32
      val = int32((xoredValue & (xoredValue-1)) ^ xoredValue)

      bucket1,bucket2 := 0,0
      for _,v := range nums {
        if v & int(val) == 0  {
            bucket1 ^= v
        } else {
            bucket2 ^= v
        }
      }

      return []int{bucket1,bucket2}
}
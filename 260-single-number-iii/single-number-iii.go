func singleNumber(nums []int) []int {

	// map based soln : TC : O(n) , SC : O(n)
	/*counts := make(map[int]int,len(nums))
	  result := make([]int,2)
	  for _,v := range nums {
	      counts[v]++
	  }

	  idx := 0
	  for val,count := range counts {
	      if count == 1 {
	        result[idx] = val
	        idx++
	      }
	  }

	  return result*/
      
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
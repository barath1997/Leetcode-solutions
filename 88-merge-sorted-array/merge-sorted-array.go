func merge(nums1 []int, m int, nums2 []int, n int)  []int{
     
	i , j := 0,0
	for j<n && i < m+n{
	   if nums1[i] == nums2[j] {
		   insert(nums1,i,nums2[j])
		   i++;j++
	   } else if nums1[i] > nums2[j] {
		   insert(nums1,i,nums2[j])
		   j++
	   } else {i++}
	}

	if j < n {
	   b := (m+n) - 1
	   for k := n-1;k>=j;k-- {
		   nums1[b] = nums2[k]
		   b--
	   }
	}

	return nums1

}

func insert(arr []int,idx,value int) {
 
   
   for i := len(arr)-2;i>=idx;i-- {
	   arr[i+1] = arr[i]
   }

   arr[idx] = value
}
func wiggleSort(nums []int)  {

   n := len(nums)
   nums1 := make([]int,n)
   sort.Ints(nums)
   idx := n-1

   for i:=1;i<n;i+=2 {  //odd
     nums1[i] = nums[idx]
     idx--
   }
   
   for i:=0;i<n;i+=2 {  //even
    nums1[i] = nums[idx]
    idx--
   }

   copy(nums,nums1) 
    
}
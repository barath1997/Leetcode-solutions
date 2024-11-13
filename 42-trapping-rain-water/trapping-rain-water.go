// Optimised soluton by precomputing lmax and rmax
func trap(height []int) int {

  n := len(height)
  if n < 2 {return 0}
  lmax,rmax := make([]int,n),make([]int,n)
  lmax[0],rmax[n-1] = height[0],height[n-1]
  res :=0

  for i:=1;i<n;i++ {
    lmax[i] = max(lmax[i-1],height[i])
  }

  for i:=n-2;i>0;i-- {
    rmax[i] = max(rmax[i+1],height[i])
  }

  for k:=1 ;k<n-1;k++ {
    res += (min(lmax[k],rmax[k]) - height[k])
  }
  
  return res
}


// TC : O(N2) , SC : O(1)
/*func trap(height []int) int {
    n := len(height)
    if n < 2 {return 0}
    result := 0
    
    leftMax,rightMax := math.MinInt,math.MinInt 
    for i:=1;i<n;i++ {
        leftMax = height[i]
        for j:=0;j<i;j++ {
         leftMax = max(leftMax,height[j])
        }
        
        rightMax = height[i]
        for j:=i+1;j<n;j++ {
            rightMax = max(rightMax,height[j])
        }
        
        result += (min(leftMax,rightMax) - height[i])
    }
    return result
}*/
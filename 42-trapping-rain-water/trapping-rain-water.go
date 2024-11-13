func trap(height []int) int {
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
}
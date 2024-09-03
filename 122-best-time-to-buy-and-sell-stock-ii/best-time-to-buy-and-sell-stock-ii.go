func maxProfit(prices []int) int {

    n := len(prices)
    if n == 1 {
        return 0
    } 
    maxProfit := 0
    
	for i:=0;i<n-1;i++ {
		if diff := prices[i+1] - prices[i]; diff > 0 { maxProfit+=diff }
	}

    return maxProfit
    
}
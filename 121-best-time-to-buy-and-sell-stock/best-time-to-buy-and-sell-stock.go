func maxProfit(prices []int) int {

    n := len(prices)
    if n == 1 {
        return 0
    } 
    maxProfit := 0
    minPrice := prices[0]

    for i:=1;i<n;i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {
            maxProfit = max(maxProfit,prices[i]-minPrice)
        }
    }
     

    return maxProfit
    
}
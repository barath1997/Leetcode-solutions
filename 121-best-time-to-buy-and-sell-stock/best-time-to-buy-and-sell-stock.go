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
            profit := prices[i] - minPrice
            if profit > maxProfit { maxProfit = profit }
        }
    }
     

    return maxProfit
    
}
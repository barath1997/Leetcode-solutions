func numDecodings(s string) int {

    n := len(s)
    if n ==0 || string(s[0]) == "0" {return 0}
    dp := make([]int,n+1)

    dp[0] = 1
    dp[1] = 1

    for i:=2;i<=n;i++ {
        if string(s[i-1]) >= "1" && string(s[i-1]) <= "9" {
            dp[i] = dp[i-1]
        }
        if string(s[i-2]) == "1" {
            dp[i] += dp[i-2]
        } else if (string(s[i-2]) == "2" && (string(s[i-1]) >= "0" && string(s[i-1]) <= "6")) {
            dp[i] += dp[i-2]
        }
    }
    return dp[n]
}
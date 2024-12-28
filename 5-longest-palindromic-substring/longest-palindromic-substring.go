// 2D DP Question, refer youtube video : youtube.com/watch?v=UflHuQj6MVA
func longestPalindrome(s string) string {

    n := len(s)

    // creating dp array
    dp := make([][]bool,n)
    for i:= range dp {
        dp[i] = make([]bool,n)
    }

    // fill diagonal places first , (i.e) single elements should be made as true
    maxlen := 1
    start := 0
    for i:=0;i<n;i++ {
        dp[i][i] = true
    }

    // go for 2 elements now
    for i:=0;i<n-1;i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = true
            start = i
            maxlen = 2
        }
    }

    // for elements greater than 2

    for k:=3;k<=n;k++ {
        for i:=0;i<n-k+1;i++ {
            j := i+k-1
            if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
                if k > maxlen {
                    maxlen = k
                    start = i
                }
            }
        }
    }

    return string(s[start:start+maxlen])
    
    
}
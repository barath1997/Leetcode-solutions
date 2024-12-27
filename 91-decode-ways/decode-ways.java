class Solution 
{
        public int numDecodings(String s) 
        {
            int n = s.length();
            if (s == null || n==0)
                return 0;
            
            int[] dp = new int[n];
            // If leading zero: NOT Contributing to Answer, hence initialised with 0
            if (s.charAt(0)=='0')
            return 0; 

            // Else, Contributing to Answer, hence Initialised with 1
            dp[0] = 1;

            int i=0;
            for (i=1; i<n; i++)
            {
                int first = Integer.valueOf(s.substring(i,i+1));
                int second = Integer.valueOf(s.substring(i-1,i+1));
        
        // Two Cases:
        // (1) If 1-9: dp[i] = dp[i-1]: 1 Entry Point
                if (first >=1 && first<=9)
                    dp[i] += dp[i-1]; // 1 Entry Point: ALWAYS BE EXECUTED

        // (2) If 10-26: dp[i] = dp[i-1] + d[i-2]: 2 Entry Point   
                if (second >=10 && second <= 26)
                {
            // dp[i] += i>=2 ? dp[i-2] : 1; //2 Entry Points
                    if (i >=2)
                        dp[i] += dp[i-2]; 
    // dp[i] = dp[i-1] + dp[i-2];
     // dp[i-1] is ALREADY CALCULATED before if part

                    else
                        dp[i] += 1;
   // if (i<1), Only 1 Digit will contribute to Answer, 
   // else, Both 2 Digits will Contribute to Answer

   // String = 26789
   // Index  = 12345    
   // At Index 2, Val = 6- Last 2 and current 6 both will contribute     
                }
            }
            
//            for (i=0; i<n; i++)
//                System.out.print("DP[i]-> " + dp[i] + " ");
            
            return dp[n-1]; // Initialised, dp[0] in beginning           
        }
}
func fib(n int) int {

    if n == 0 {
        return 0
    }
    
   fibNumbers := make([]int,n+1)

   first,second := 0,1
   fibNumbers[0],fibNumbers[1] = first,second
   for i:=2;i<=n;i++ {
    fibNumbers[i] = fibNumbers[i-1] + fibNumbers[i-2]
   }
   
   return fibNumbers[n]
}
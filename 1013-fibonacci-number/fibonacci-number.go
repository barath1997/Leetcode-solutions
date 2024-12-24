// usin maths
func fib(n int) int {
  
  phi := float64((math.Sqrt(5) + 1) / 2)

  return int(math.Round((math.Pow(phi,float64(n)) / math.Sqrt(5))))

} 

// using DP
/*func fib(n int) int {

   if n ==0 {return 0} 
   
   res := make([]int,n+1)
   if  n <= 2 {return 1}
   res[1] = 1
   res[2] = 1

   for i:=3;i<=n;i++ {
    res[i] = res[i-1] + res[i-2]
   } 
   
   return res[n]
}*/

// using iterative approach
/*func fib(n int) int {

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
}*/

// using recursive approach
/*func fib(n int) int {
   
    if n == 1 || n ==0 {
        return n
    }
    return fib(n-1) + fib(n-2)

}*/
func isPowerOfTwo(n int) bool {
    
   if n < 1 {return false}
   if n == 1 { return true }

   if n % 2 == 0 {
      
    if n & (n-1) == 0 {return true}

   }

   return false


}
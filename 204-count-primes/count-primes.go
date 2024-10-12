func countPrimes(n1 int) int {

    if n1 <= 2 {return 0}
    if n1 == 3 {return 1}
    
   
   count :=0

   primes := make([]bool,n1+1)

   for i:=0;i<=n1;i++ {
    primes[i] = true
   }

   for p := 2;p*p<=n1;p++ {
     if primes[p] == true {
        for j:=p*p;j<=n1;j+=p {
            primes[j] = false
        }
     }
   }

   for i:=2;i<n1;i++ {
    if primes[i] == true {
        count++
    }
   }
   
   return count 
}


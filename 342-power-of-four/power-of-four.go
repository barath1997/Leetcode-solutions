func isPowerOfFour(n int) bool {
    
    if n == 1 { return true}
    if n%2 != 0 || n == 0 {return false}
    
    if n%4 == 0 {
       return isPowerOfFour(n/4) 
    } else {
        return false
    }

}
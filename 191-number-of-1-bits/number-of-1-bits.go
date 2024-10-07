func hammingWeight(n int) int {
    
    count := 0
    for n != 0 { 
    if n & 1 == 1 { 
        count++ 
    }
     n >>= 1   // right shift 10>>1 is 5 and 5>>1 is 2 and 2>>1 = 1 , n keeps diving 2 on right shift and        the setbits keep decreasing on every rightshift, if all setbits become 0 then n becomes 0 
    }

    return count
}
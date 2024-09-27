func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {return true}
    l := len(flowerbed)
    
    for i:=0;i<l;i++ {
    if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
       flowerbed[i] = 1;
       n--;
       if n == 0{
        return true
       }
    } 
    }
    return false 
}
func maxProduct(nums []int) int {

   var maxProduct int64 = math.MinInt
   var currProduct int = 1
    for _,val := range nums { //prefixproduct
        currProduct = val*currProduct
        maxProduct = max(maxProduct, int64(currProduct))

        if currProduct == 0 {
            currProduct = 1
        }
    }

    currProduct = 1

    for i:=len(nums)-1;i>=0;i-- {
        currProduct = currProduct * nums[i]
        maxProduct = max(maxProduct,int64(currProduct))

        if currProduct == 0 {
            currProduct = 1
        }

    }

    return int(maxProduct)
}
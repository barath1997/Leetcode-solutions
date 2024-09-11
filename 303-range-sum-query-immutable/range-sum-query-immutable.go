type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    n := NumArray{nums : make([]int,len(nums))}

    for i:=1;i<len(nums);i++ {
        nums[i] = nums[i] + nums[i-1]
    }

    n.nums = nums
    return n
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 && right == 0{return this.nums[0]}
    if left == 0 {
        return this.nums[right]
    }

    if left !=0 {
        return this.nums[right] - this.nums[left-1]
    }

    return -1
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
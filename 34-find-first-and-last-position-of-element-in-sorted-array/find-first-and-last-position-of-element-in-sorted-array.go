func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1, -1}
    }
    
    minIndex := findMinIndex(0, n - 1, target, nums)
    maxIndex := findMaxIndex(0, n - 1, target, nums)
    
    return []int{minIndex, maxIndex}
}

func findMinIndex(low, high, target int, arr []int) int {
    if high >= low {
        mid := low + (high - low) / 2
        
        if arr[mid] == target && (mid == 0 || arr[mid-1] < target) {
            return mid
        } else if arr[mid] >= target {
            return findMinIndex(low, mid - 1, target, arr)
        } else {
            return findMinIndex(mid + 1, high, target, arr)
        }
    }
    return -1
}

func findMaxIndex(low, high, target int, arr []int) int {
    if high >= low {
        mid := low + (high - low) / 2
        
        if arr[mid] == target && (mid == len(arr) - 1 || arr[mid + 1] > target) {
            return mid
        } else if arr[mid] <= target {
            return findMaxIndex(mid + 1, high, target, arr)
        } else {
            return findMaxIndex(low, mid - 1, target, arr)
        }
    }
    return -1
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || len(nums) < 2 {
		return false
	}

	bucket := make(map[int]int)
	w := max(valueDiff , 1) // Width of the bucket

	for i, num := range nums {
		bucketKey := num / w

		// Check the current bucket
		if val, exists := bucket[bucketKey]; exists && abs(num-val) <= valueDiff {
			return true
		}

		// Check the left neighbor bucket
		if val, exists := bucket[bucketKey-1]; exists && abs(num-val) <= valueDiff {
			return true
		}

		// Check the right neighbor bucket
		if val, exists := bucket[bucketKey+1]; exists && abs(num-val) <= valueDiff {
			return true
		}

		// Add the current number to the bucket
		bucket[bucketKey] = num

		// Remove the element that is out of the indexDiff range
		if i >= indexDiff {
			delete(bucket, nums[i-indexDiff]/w)
		}
	}

	return false
}

// Helper function to compute the absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

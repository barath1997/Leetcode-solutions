// https://www.youtube.com/watch?v=OyZFFqQtu98

func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{} // Initialize result as a slice of slices
    ds := []int{}       // This will hold the current combination
    return findCombinations(0, target, candidates, result, ds)
}

func findCombinations(index, target int, inputArr []int, result [][]int, ds []int) [][]int {
    // Base case: if target is 0, we found a valid combination
    if target == 0 {
        // Append a copy of ds to result
        combination := make([]int, len(ds))
        copy(combination, ds)
        result = append(result, combination)
        return result
    }

    // If we exceed the index or target, return
    if index >= len(inputArr) || target < 0 {
        return result
    }

    // Include the current candidate
    if inputArr[index] <= target {
        ds = append(ds, inputArr[index])
        result = findCombinations(index, target-inputArr[index], inputArr, result, ds)
        ds = ds[:len(ds)-1] // Backtrack
    }

    // Exclude the current candidate and move to the next
    result = findCombinations(index+1, target, inputArr, result, ds)

    return result
}
package arrayshashing

func containsDuplicate(nums []int) bool {
	// first declare a map
	numMap := make(map[int]bool)
	for _, num := range nums {
		if numMap[num] {
			return true
		}
		numMap[num] = true
	}
	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
//another solution

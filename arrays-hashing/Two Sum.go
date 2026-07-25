package arrayshashing

import "fmt"

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	for i, num := range nums {
		NumToFind := target - num
		if j, ok := valueMap[NumToFind]; ok {
			return []int{i, j}
		}
		valueMap[NumToFind] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result) // [0 1]
}

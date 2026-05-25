package main

import "fmt"

func findMinMax(nums []int) (int, int) {
	// Handle empty array case
	if len(nums) == 0 {
		return 0, 0
	}

	// Initialize min and max with the first element
	min := nums[0]
	max := nums[0]

	// Start loop from the second element
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return min, max
}

func main() {
	arr := []int{3, 5, 1, 9, 2, -4, 8}
	min, max := findMinMax(arr)
	
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

}

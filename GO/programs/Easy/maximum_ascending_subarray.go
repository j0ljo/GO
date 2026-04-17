// 1800. Maximum Ascending subarray sum 

func maxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// Still ascending! Keep adding.
			currentSum += nums[i]
		} else {
			// Streak broken. Reset current sum to the current number.
			currentSum = nums[i]
		}

		// Update the record if the current streak is the biggest we've seen.
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

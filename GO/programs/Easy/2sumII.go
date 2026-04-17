// 167. Two Sum II 
//
//
func twoSum(numbers []int, target int) []int {
	i := 0 
	j := len(numbers) - 1 

	for i < j {
		currenSum := numbers[i] + numbers[j]

		if currenSum == target {
			return []int{i+1, j+1}

		}
		if currenSum < target {
			i++
		} else {			// this is for currenSum > target 
			j-- 
		}
	}
	return []int{} 
}


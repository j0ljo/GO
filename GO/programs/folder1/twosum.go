
func twoSum(nums []int, target int) []int {
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++ {
			if nums[i] + nums[j] == target {
				return [] int{i,j}
			}
		}

	}
	return nil 
}


// SOlve using Hash maps for faster run time 

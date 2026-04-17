// 15. 3Sum 
//
//
func threeSum(nums []int) [][]int {
	n := len(nums) + 1

	for i:=0; i<n ; i++ {
		for j:=0; j<n ; j++ {
			for k:=0; k<n; k++ {
				if nums[i] + nums[j] + nums[k] == 0 && i!=j!=k {
					return []int{i,j,k}
				}

			}
		}
	}
	return []int{}

}


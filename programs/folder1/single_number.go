// 136. Single Number 
// Use XOR (efficient) or use Maps 
//
//


func singleNumber(nums []int) int {
	counts := make(map[int]int)

	for _, n := range nums {
		counts[n] ++ 
	}
	for n, count := range counts {
		if count == 1 {
			return n 
		}
	}
	return 0 
}

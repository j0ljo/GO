// 3783. Mirror Distance of an integer 
//

func mirrorDistance(n int) int {
	// 1. Convert int to string
	s := strconv.Itoa(n)
	runes := []rune(s)
	
	// 2. Reverse the runes (digits)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// 3. Convert reversed string back to int
	reversedInt, _ := strconv.Atoi(string(runes))
	
	// 4. Return absolute difference
	diff := n - reversedInt
	if diff < 0 {
		return -diff
	}
	return diff
}

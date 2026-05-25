func isAnagram(s string, t string) bool {
	if len(s) != len(t) { 
		return false 
	}

	counts := make(map [rune]int )

	for i, charS := range s {
		counts[charS]++
		counts[rune(t[i])]-- 

	}
	for _, count := range counts{
		if count != 0 {
			return false 
		}
	}
	return true



}

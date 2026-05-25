
func romanToInt(s string) int {
	mp := map[byte]int {
	
		'I':1,   // change all the "" to '' for rune/int/byte 
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	
	}
	
	total := 0 
	n := len(s) 

	for i:=0 ; i<n ; i++ {
		currenVal := mp[s[i]]

		if i<n-1 && currenVal < mp[s[i+1]] {

			total -= currenVal
		} else {
			total += currenVal
		}
	}
	return total 
	
} 


// 917. Reverse only letters 
//

func reverseOnlyLetters(s string) string {
	r := []rune(s) 
	i, j := 0, len(r) - 1 
	for i<j {
		if !unicode.IsLetter(r[i]) {
			i++
			continue 
		}
		if !unicode.IsLetter(r[j]) {
			j--
			continue 
		}
		r[i], r[j] = r[j], r[i]
		i++
		j-- 
	}
	return string(r) 
}

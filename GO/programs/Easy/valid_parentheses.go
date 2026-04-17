// 20. Valid Parentheses 
//
// stack = append(stack, char) ; you have to manually define the top of the stack 
//
func isValid(s string) bool {

	stack := []rune{}
	// Push 
	for _, char := range s {


		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)

			continue
		}

	if len(stack) == 0 {
		return false 
	}
	// Pop 
	top := stack[len(stack) - 1]
	stack = stack[:len(stack)-1]

	if (char == ')' && top != '(') ||
		 (char == '}' && top != '{') ||
		 (char == ']' && top != '[' ) {
			 return false 
		 }

	
	}
	return len(stack) == 0 
}

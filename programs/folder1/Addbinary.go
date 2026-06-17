
package main

import (
	"fmt"

	// "strconv"
)

// The entry point of your program
func main() {
	result := addBinary("1010", "1011")
	fmt.Println("Result:", result) // Expected: 10101
}

func addBinary(a string, b string) string {
	var result []byte
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = append(result, byte(sum%2+'0'))
		carry = sum / 2
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}

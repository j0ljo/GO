package main 

import "fmt"

func sum(num1 int, num2 int) (int) {
	a := num1
	b := num2
	result := a + b
	return result 

}

func main() {
	result := sum(10,25)
	fmt.Printf("Sum is: ", result)
}

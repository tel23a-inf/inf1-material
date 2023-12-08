package main

import "fmt"

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}
	return result
}

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(Factorial(3)))

	n1 := 3
	n2 := Factorial(n1)
	n3 := Factorial(n2)
	fmt.Println(n3)

	fmt.Println(Factorial(20))
	fmt.Println(Factorial(21))
}

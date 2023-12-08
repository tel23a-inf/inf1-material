package main

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(Factorial(3)))

	n1 := 3
	n2 := Factorial(n1)
	n3 := Factorial(n2)
	fmt.Println(n3)

	fmt.Println(Factorial(20))

	// Output:
	// 120
	// 1
	// 720
	// 720
	// 2432902008176640000
}

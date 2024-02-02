package rekursion

import "fmt"

func ExampleCountDown() {
	CountDown(3)

	// Output:
}

func ExampleCountUpNegative() {
	CountUpNegative(-3)

	// Output:
}

func ExampleCountUp() {
	CountUp(3)

	// Output:
}

func ExampleFactorial() {
	fmt.Println(Factorial(3))

	// Output:
}

func ExampleAckermann() {
	//fmt.Println(Ackermann(1, 0))
	//fmt.Println(Ackermann(2, 0))
	//fmt.Println(Ackermann(1, 1))
	//fmt.Println(Ackermann(2, 1))
	//fmt.Println(Ackermann(3, 1))
	//fmt.Println(Ackermann(4, 1))
	//fmt.Println(Ackermann(3, 3))
	fmt.Println(Ackermann(4, 4))

	// Output:
}

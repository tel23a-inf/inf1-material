package rekursion

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

func CountUpNegative(n int) {
	// Zusätzliche Variable?
	if n >= 0 { // ?
		return
	}
	fmt.Println(n)
	CountUpNegative(n + 1)
}

func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := n * Factorial(n-1)
	return result
}

func Ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	}
	//	if m > 0 && n == 0 {
	if n == 0 {
		return Ackermann(m-1, 1)
	}
	return Ackermann(m-1, Ackermann(m, n-1))
}

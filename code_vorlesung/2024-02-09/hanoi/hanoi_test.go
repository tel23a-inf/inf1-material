package hanoi

import "fmt"

func ExampleHanoi() {
	fmt.Println(Hanoi(1, "A", "B", "C"))
	fmt.Println()
	fmt.Println(Hanoi(2, "A", "B", "C"))
	fmt.Println()
	fmt.Println(Hanoi(3, "A", "B", "C"))

	// Output:
	// A -> C
	// 1
	//
	// A -> B
	// A -> C
	// B -> C
	// 3
	//
	// A -> C
	// A -> B
	// C -> B
	// A -> C
	// B -> A
	// B -> C
	// A -> C
	// 7
}

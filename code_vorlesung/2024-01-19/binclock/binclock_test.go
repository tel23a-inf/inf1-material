package binclock

import "fmt"

func ExampleToBinary() {
	fmt.Println(ToBinary(1))
	fmt.Println(ToBinary(3))
	fmt.Println(ToBinary(12))
	fmt.Println(ToBinary(59))
	fmt.Println(ToBinary(60))

	// Output:
	// [0 0 0 0 0 1]
	// [0 0 0 0 1 1]
	// [0 0 1 1 0 0]
	// [1 1 1 0 1 1]
	// [1 1 1 1 0 0]
}

func ExampleBinaryToString() {
	fmt.Println(BinaryToString([]int{0, 0, 0, 0, 0, 1}))
	fmt.Println(BinaryToString([]int{0, 0, 0, 0, 1, 1}))
	fmt.Println(BinaryToString([]int{0, 0, 1, 1, 0, 0}))
	fmt.Println(BinaryToString([]int{1, 1, 1, 0, 1, 1}))
	fmt.Println(BinaryToString([]int{1, 1, 1, 1, 0, 0}))

	// Output:
	// |     *|
	// |    **|
	// |  **  |
	// |*** **|
	// |****  |
}

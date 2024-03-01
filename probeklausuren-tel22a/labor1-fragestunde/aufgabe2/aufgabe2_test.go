package aufgabe2

import "fmt"

func ExampleColour() {
	fmt.Println(Colour(0))
	fmt.Println(Colour(9))
	fmt.Println(Colour(10))
	fmt.Println(Colour(20))
	fmt.Println(Colour(30))
	fmt.Println(Colour(40))
	fmt.Println(Colour(42))
	fmt.Println(Colour(52))

	// Output:
	// Rot
	// Rot
	// Rot-Gelb
	// Grün
	// Gelb
	// Rot
	// Rot
	// Rot-Gelb
}

func ExampleColour2() {
	fmt.Println(Colour2(0))
	fmt.Println(Colour2(9))
	fmt.Println(Colour2(10))
	fmt.Println(Colour2(20))
	fmt.Println(Colour2(30))
	fmt.Println(Colour2(40))
	fmt.Println(Colour2(42))
	fmt.Println(Colour2(52))

	// Output:
	// Rot
	// Rot
	// Rot-Gelb
	// Grün
	// Gelb
	// Rot
	// Rot
	// Rot-Gelb
}

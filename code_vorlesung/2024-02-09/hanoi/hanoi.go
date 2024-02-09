package hanoi

import "fmt"

// func Hanoi(h int, s, m, z string) int {
// 	if h == 1 {
// 		return Move(s, z)
// 	} else {
// 		counter := 0
// 		counter += Hanoi(h-1, s, z, m)
// 		counter += Move(s, z)
// 		counter += Hanoi(h-1, m, s, z)
// 		return counter
// 	}
// }

// func Hanoi(h int, s, m, z string) int {
// 	counter := 0
// 	if h > 0 {
// 		counter += Hanoi(h-1, s, z, m)
// 		counter += Move(s, z)
// 		counter += Hanoi(h-1, m, s, z)
// 	}
// 	return counter
// }

func Hanoi(h int, s, m, z string) int {
	if h == 0 {
		return 0
	}

	counter := 0
	counter += Hanoi(h-1, s, z, m)
	counter += Move(s, z)
	counter += Hanoi(h-1, m, s, z)

	return counter
}

func Move(s, z string) int {
	fmt.Printf("%v -> %v\n", s, z)
	return 1
}

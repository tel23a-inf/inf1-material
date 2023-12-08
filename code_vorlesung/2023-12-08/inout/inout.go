package main

import "fmt"

func main() {
	fmt.Println("Bitte eine Zahl eingeben:")

	var x int

	fmt.Scanln(&x)

	fmt.Println("Sie haben folgende Zahl eingegeben:", x)

	if x != 42 {
		fmt.Println("Die Zahl war falsch!")
	} else {
		fmt.Println("Gute Wahl!")
	}
}

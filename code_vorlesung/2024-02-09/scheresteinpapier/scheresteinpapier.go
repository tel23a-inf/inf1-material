package main

import "fmt"

func main() {
	x := ReadInput()
	fmt.Println(x)
}

func ReadInput() int {
	fmt.Println("Bitte wähle:")
	fmt.Println("1: Schere")
	fmt.Println("2: Stein")
	fmt.Println("3: Papier")

	var input string
	//fmt.Print("Adresse von input: ")
	//fmt.Println(&input)

	fmt.Scanln(&input)

	result := 0

	if len(input) == 1 {
		result = int(input[0]) - '0'

		if result <= 3 && result >= 1 {
			return result
		}
	}
	fmt.Println("Das war nix. Versuch es noch einmal!")
	return ReadInput()
}

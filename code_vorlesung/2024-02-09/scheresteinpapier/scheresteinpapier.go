package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Run()
}

func Run() {
	h := GetHumanInput()
	if h == 0 {
		fmt.Println("Auf Wiedersehen, bis zum nächsten Mal.")
		return
	}
	PrintResult(CompareInput(h, GetBotInput()))
	Run()
}

func PrintResult(result int) {
	switch result {
	case 0:
		fmt.Println("Unentschieden")
	case 1:
		fmt.Println("Der Mensch gewinnt")
	case 2:
		fmt.Println("Der Bot gewinnt")
	}
}

// Erwartet die Eingaben beider Spieler, liefert:
// 0: unentschieden
// 1: human gewinnt
// 2: bot gewinnt
func CompareInput(human, bot int) int {
	if bot == human {
		return 0
	}
	if bot == 1 && human == 2 || bot == 2 && human == 3 || bot == 3 && human == 1 {
		return 1
	}
	return 2
}

func GetBotInput() int {
	choice := rand.Int()%3 + 1
	fmt.Printf("Der Bot wählt %d.\n", choice)

	/* Alternative:
	fmt.Print("Der Bot wählt ")
	fmt.Print(choice)
	fmt.Println(".")
	*/
	return choice
}

func GetHumanInput() int {
	fmt.Println("Bitte wähle:")
	fmt.Println("1: Schere")
	fmt.Println("2: Stein")
	fmt.Println("3: Papier")
	fmt.Println("0: Beenden")

	var input string
	//fmt.Print("Adresse von input: ")
	//fmt.Println(&input)

	fmt.Scanln(&input)

	result := 0

	if len(input) == 1 {
		result = int(input[0]) - '0'

		if result <= 3 && result >= 0 {
			return result
		}
	}
	fmt.Println("Das war nix. Versuch es noch einmal!")
	return GetHumanInput()
}

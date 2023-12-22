package listen

import (
	"fmt"
	"strings"
)

func IntLists() {
	// Erzeugen einer Liste mit Elementen vom Typ int.
	var list1 []int
	fmt.Println(list1)

	// Anhängen von Elementen an die Liste.
	list1 = append(list1, 42, 23, 1000000)
	fmt.Println(list1)

	// Zugreifen auf einzelne Elemente der Liste.
	fmt.Println(list1[0])
	fmt.Println(list1[1])
	el := list1[1] + 25
	fmt.Println(el)

	// Ändern eines Listenelements.
	list1[1] = 103
	fmt.Println(list1[1])
}

func StringLists() {
	// Erzeugen einer Liste mit Elementen vom Typ string.
	var list2 []string
	fmt.Println(list2)

	// Anhängen von Elementen an die Liste.
	list2 = append(list2, "Hallo", "Welt")
	fmt.Println(list2)

	// Erzeugen einer Liste mit Werten.
	list3 := []string{"Foo", "Bar"}
	fmt.Println(list3)
}

// Liefert eine Liste mit dem zehnfachen der Zahlen von 1 bis n.
func NumberList(n int) []int {
	result := []int{} // Leere Liste.
	for i := 1; i <= n; i++ {
		result = append(result, 10*i)
	}
	return result
}

// Erwartet eine Zahl n und liefert eine Liste mit den Ziffern von n.
func Digits(n int) []int {
	result := []int{}

	// Schleife, die die einzelnen Ziffern aus der Zahl
	// berechnet und in result schreibt.
	for n != 0 {
		result = append(result, n%10) // z.B. 321 % 10 == 1
		n = n / 10
	}

	// Liste umdrehen (generiert von VsCode).
	// for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
	// 	result[i], result[j] = result[j], result[i]
	// }
	result = Revert(result)

	return result
}

// Erwartet eine Liste und liefert eine Liste mit den Einträgen in umgekehrter Reihenfolge.
func Revert(list []int) []int {
	result := []int{}

	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}

	return result
}

func PrintRow(row []string) {
	fmt.Print("|")
	for i := 0; i < len(row); i++ {
		fmt.Print(" ")
		fmt.Print(row[i])
		fmt.Print(" |")
	}
	fmt.Println()
}

func PrintRow2(row []string) {
	fmt.Print("|")
	for _, el := range row {
		fmt.Print(" ")
		fmt.Print(el)
		fmt.Print(" |")
	}
	fmt.Println()
}

func PrintBoard(board [][]string) {
	// "+---" * len(board[0])
	border := Border(len(board[0]))
	fmt.Println(border)
	for _, row := range board {
		PrintRow(row)
		fmt.Println(border)
	}
}

func Border(width int) string {
	return fmt.Sprintf("%s+", strings.Repeat("+---", width))

	// Alternative ohne Printf
	// return strings.Repeat("+---", width) + "+"

	// "händische" Alternative
	// border := ""
	// for i := 0; i < width; i++ {
	// 	border += "+---"
	// }
	// border += "+"
	// return border
}

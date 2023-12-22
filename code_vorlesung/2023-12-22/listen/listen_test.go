package listen

import "fmt"

func ExampleIntLists() {
	IntLists()

	// Output:
	// []
	// [42 23 1000000]
	// 42
	// 23
	// 48
	// 103
}

func ExampleStringLists() {
	StringLists()

	// Output:
	// []
	// [Hallo Welt]
	// [Foo Bar]
}

func ExampleNumberList() {
	list1 := NumberList(15)
	fmt.Println(list1)

	second := list1[1]
	fmt.Println(second)

	// Output:
}

func ExampleDigits() {
	fmt.Println(Digits(42))
	fmt.Println(Digits(321))
	fmt.Println(Digits(5))
	fmt.Println(Digits(0))

	// Output:
	// [4 2]
	// [3 2 1]
	// [5]
	// []
}

func ExamplePrintRow() {
	row1 := []string{"1", "2", "3"}
	PrintRow(row1)
	PrintRow2(row1)

	row2 := []string{" ", "X", " ", "O", "O"}
	PrintRow(row2)
	PrintRow2(row2)

	// Output:
	// | 1 | 2 | 3 |
	// | 1 | 2 | 3 |
	// |   | X |   | O | O |
	// |   | X |   | O | O |
}

func ExamplePrintBoard() {
	// Ein Spielfeld ist eine Liste,
	// deren Elemente Listen von Strings sind.
	// Jede solche innere Liste ist eine Zeile auf
	// dem Spielfeld.
	b1 := [][]string{
		{" ", "X", " ", "X"},
		{" ", "X", "O", "O"},
		{"O", " ", " ", "X"},
	}

	PrintBoard(b1)

	// Wie muss eine Funktionssignatur
	// plus Gerüst für PrintBoard aussehen?

	// Output:
	// +---+---+---+---+
	// |   | X |   | X |
	// +---+---+---+---+
	// |   | X | O | O |
	// +---+---+---+---+
	// | O |   |   | X |
	// +---+---+---+---+
}

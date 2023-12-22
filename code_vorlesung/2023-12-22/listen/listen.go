package listen

import "fmt"

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
}

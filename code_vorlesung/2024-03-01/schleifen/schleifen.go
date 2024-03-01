package schleifen

import "fmt"

func Range1() {
	l1 := []string{"A", "B", "C", "D", "E"}
	for i := range l1 {
		fmt.Println(i)
	}
}

func Range2() {
	l1 := []string{"A", "B", "C", "D", "E"}
	for _, i := range l1 {
		fmt.Println(i)
	}
}

func Range3() {
	l1 := []string{"A", "B", "C", "D", "E"}
	for _, element := range l1 {
		fmt.Println(element)
	}
}

func Range4() {
	l1 := []string{"A", "B", "C", "D", "E"}
	for i, v := range l1 {
		fmt.Printf("%d: %s\n", i, v)
	}
}

func Range5() {
	l1 := []string{"A", "B", "C", "D", "E"}
	fmt.Print("[")
	for _, v := range l1[:len(l1)-1] {
		fmt.Printf("%s,", v)
	}
	//fmt.Print(l1[len(l1)-1], "]")
	//fmt.Print(l1[len(l1)-1] + "]")
	fmt.Printf("%s]", l1[len(l1)-1])
}

func Range6() {
	l1 := []string{"A", "B", "C", "D", "E"}

	contents := ""
	for _, v := range l1[:len(l1)-1] {
		contents += fmt.Sprintf("%s,", v)
	}
	contents += l1[len(l1)-1]

	fmt.Printf("[%s]", contents)
}

func Range7() {
	l1 := []string{"A", "B", "C", "D", "E"}

	contents := ""
	for _, v := range l1 {
		contents += fmt.Sprintf("%s,", v)
	}

	fmt.Printf("[%s]", contents[:len(contents)-1])
}

func Range8() {
	l1 := []string{"A", "B", "C", "D", "E"}

	contents := ""
	for i := 0; i < len(l1); i++ {
		contents += fmt.Sprintf("%s,", l1[i])
	}

	fmt.Printf("[%s]", contents[:len(contents)-1])
}

func Range9() string {
	l1 := []string{"A", "B", "C", "D", "E"}

	contents := ""
	for _, v := range l1 {
		contents += fmt.Sprintf("%s,", v)
	}

	return fmt.Sprintf("[%s]", contents[:len(contents)-1])
}

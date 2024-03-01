package aufgabe7

import "fmt"

func ExampleIntersect() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}
	l3 := []int{}
	l4 := []int{3, 13, 11, 9, 1, 7, 5}

	fmt.Println(Intersect(l1, l2))
	fmt.Println(Intersect(l1, l3))
	fmt.Println(Intersect(l3, l2))
	fmt.Println(Intersect(l3, l3))
	fmt.Println(Intersect(l1, l4))

	// Output:
	// [1 7 13]
	// []
	// []
	// []
	// [1 3 5 7 9 11 13]
}

func ExampleIntersectSorted() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}
	l3 := []int{}

	fmt.Println(IntersectSorted(l1, l2))
	fmt.Println(IntersectSorted(l1, l3))
	fmt.Println(IntersectSorted(l3, l2))
	fmt.Println(IntersectSorted(l3, l3))

	// Output:
	// [1 7 13]
	// []
	// []
	// []
}

func ExampleIntersectSortedRec() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}
	l3 := []int{}

	fmt.Println(IntersectSortedRec(l1, l2))
	fmt.Println(IntersectSortedRec(l1, l3))
	fmt.Println(IntersectSortedRec(l3, l2))
	fmt.Println(IntersectSortedRec(l3, l3))

	// Output:
	// [1 7 13]
	// []
	// []
	// []
}

func ExampleIntersectSortedRec2() {
	l1 := []int{1, 3, 5, 7, 9, 11, 13}
	l2 := []int{1, 4, 7, 10, 13}
	l3 := []int{}

	fmt.Println(IntersectSortedRec2(l1, l2))
	fmt.Println(IntersectSortedRec2(l1, l3))
	fmt.Println(IntersectSortedRec2(l3, l2))
	fmt.Println(IntersectSortedRec2(l3, l3))

	// Output:
	// [1 7 13]
	// []
	// []
	// []
}

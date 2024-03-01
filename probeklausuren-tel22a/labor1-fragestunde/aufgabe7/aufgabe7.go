package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Intersect.
PUNKTE: 10
BEWERTUNG:
*/

// Intersect erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen, die sowohl in l1 als auch
// in l2 vorhanden sind.
// Hinweis: Gehen Sie davon aus, dass jedes Element pro Liste höchstens einmal vorkommt.
func Intersect(l1, l2 []int) []int {
	result := []int{}

	for _, v1 := range l1 {
		for _, v2 := range l2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}

	return result
}

func IntersectSorted(l1, l2 []int) []int {
	result := []int{}

	for p1, p2 := 0, 0; p1 < len(l1) && p2 < len(l2); {
		if l1[p1] == l2[p2] {
			result = append(result, l1[p1])
		}
		if l1[p1] < l2[p2] {
			p1++
		} else {
			p2++
		}
	}
	return result
}

func IntersectSortedRec(l1, l2 []int) []int {
	result := []int{}

	if len(l1) > 0 && len(l2) > 0 {
		if l1[0] == l2[0] {
			result = append(result, l1[0])
			return append(result, IntersectSortedRec(l1[1:], l2[1:])...)
		}
		if l1[0] > l2[0] {
			return IntersectSortedRec(l1, l2[1:])
		}
		if l1[0] < l2[0] {
			return IntersectSortedRec(l1[1:], l2)
		}
	}

	return result
}

func IntersectSortedRec2(l1, l2 []int) []int {

	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	if l1[0] == l2[0] {
		return append([]int{l1[0]}, IntersectSortedRec2(l1[1:], l2[1:])...)
	}
	if l1[0] > l2[0] {
		return IntersectSortedRec2(l1, l2[1:])
	}
	return IntersectSortedRec2(l1[1:], l2)
}

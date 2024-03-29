package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ArraySums.
PUNKTE: 15
BEWERTUNG:
*/

// ArraySums erwartet eine int-Slice l erwartet und liefert eine int-Slice,
// die an Stelle n die Summe der Elemente aus l bis zu Stelle n enthält.
func ArraySums(l []int) []int {
	// BEGIN-SOLUTION
	result := []int{}
	sum := 0
	for _, val := range l {
		sum += val
		result = append(result, sum)
	}
	return result
	// END-SOLUTION
}

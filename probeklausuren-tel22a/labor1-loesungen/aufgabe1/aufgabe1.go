package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MultFive.
PUNKTE: 5
BEWERTUNG:
*/

// MultFive erwartet eine Liste von Zahlen und liefert
// das Produkt der Elemente an den durch 5 teilbaren Positionen.
func MultFive(list []int) int {
	// BEGIN-SOLUTION
	result := 1
	for pos, val := range list {
		if pos%5 == 0 {
			result *= val
		}
	}
	return result
	// END-SOLUTION
}

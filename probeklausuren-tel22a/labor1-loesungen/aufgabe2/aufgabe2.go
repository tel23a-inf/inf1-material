package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Colour.
PUNKTE: 10
BEWERTUNG:
*/

// Colour erwartet eine Zahl t, die als Sekundenangabe aufgefasst werden soll.
// Die Funktion liefert einen String, der angibt, welche Farbe eine Ampel nach t Sekunden hat.
// Dabei soll die Ampel bei Rot starten und alle 10 Sekunden umschalten.
func Colour(t int) string {
	// BEGIN-SOLUTION
	phase_count := t / 10
	phase := phase_count % 4
	outputs := []string{"Rot", "Rot-Gelb", "Grün", "Gelb"}
	return outputs[phase]
	// END-SOLUTION
}

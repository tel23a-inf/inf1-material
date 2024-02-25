package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Zip.
Die Funktion muss rekursiv sein.
PUNKTE: 15
BEWERTUNG:
*/

// Zip(), erwartet zwei Strings und mischt sie zusammen,indem Sie
// immer abwechselnd ein Zeichen aus dem einen und dem anderen nimmt.
// Sind die Strings nicht gleich lang, soll der Rest des längeren ans
// Ende gehängt werden.
func Zip(s1, s2 string) string {
	// BEGIN-SOLUTION
	if len(s1) == 0 {
		return s2
	}
	if len(s2) == 0 {
		return s1
	}
	return s1[:1] + s2[:1] + Zip(s1[1:], s2[1:])
	// END-SOLUTION
}

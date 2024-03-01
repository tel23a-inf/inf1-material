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
	t = t % 40
	if t < 10 {
		return "Rot"
	}
	if t < 20 {
		return "Rot-Gelb"
	}
	if t < 30 {
		return "Grün"
	}
	if t < 40 {
		return "Gelb"
	}
	return ""
}

func Colour2(t int) string {
	return []string{"Rot", "Rot-Gelb", "Grün", "Gelb"}[(t/10)%4]
}

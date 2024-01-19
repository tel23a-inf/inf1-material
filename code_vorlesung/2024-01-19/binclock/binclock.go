package binclock

import "time"

// ToBinary erwartet eine Zahl 0 <= n <= 60 und liefert eine Liste von 0 oder 1.
func ToBinary(n int) []int {
	result := []int{0, 0, 0, 0, 0, 0}
	result[5] = n % 2
	n = n / 2
	result[4] = n % 2
	n = n / 2
	result[3] = n % 2
	n = n / 2
	result[2] = n % 2
	n = n / 2
	result[1] = n % 2
	n = n / 2
	result[0] = n % 2
	return result
}

// BinaryToString erwartet eine Liste aus 0 oder 1 und liefert einen String aus "*" und Leerzeichen.
func BinaryToString(bindigits []int) string {
	result := ""
	for _, d := range bindigits {
		if d == 1 {
			result += "*"
		} else {
			result += " "
		}
	}
	return "|" + result + "|"
}

// UnixTime liefert die Uhrzeit als int (Sekunden seit 1.1.1970).
func UnixTime() int {
	return int(time.Now().Unix())
}

// Seconds liefert die Sekunden der aktuellen Minute.
func Seconds(unix int) int {
	return unix % 60
}

// Minutes liefert die Minuten der aktuellen Stunde.
func Minutes(unix int) int {
	return (unix / 60) % 60
}

// Hours liefert die Stunden des aktuellen Tages.
func Hours(unix int) int {
	return (unix / (60 * 60)) % 24
}

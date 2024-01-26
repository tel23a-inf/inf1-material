package numbers

import "slices"

// DigitList returns a list of digits of the number n.
// The result starts with the most significant digit.
func DigitList(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Reverse the slice using the slices package.
	slices.Reverse(digits)

	return digits
}

// DigitToString100 returns a german string representation of the
// number n as it would appear in the 100 digit.
func DigitToString100(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return "einhundert"
	case 2:
		return "zweihundert"
	case 3:
		return "dreihundert"
	case 4:
		return "vierhundert"
	case 5:
		return "fünfhundert"
	case 6:
		return "sechshundert"
	case 7:
		return "siebenhundert"
	case 8:
		return "achthundert"
	case 9:
		return "neunhundert"
	default:
		panic("DigitToString100: invalid digit")
	}
}

// DigitToString10 returns a german string representation of the
// number n as it would appear in the 10 digit.
func DigitToString10(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return "zehn"
	case 2:
		return "zwanzig"
	case 3:
		return "dreißig"
	case 4:
		return "vierzig"
	case 5:
		return "fünfzig"
	case 6:
		return "sechzig"
	case 7:
		return "siebzig"
	case 8:
		return "achtzig"
	case 9:
		return "neunzig"
	default:
		panic("DigitToString10: invalid digit")
	}
}

// DigitToString1 returns a german string representation of the
// number n as it would appear in the 1 digit.
func DigitToString1(n int) string {

	dstrings := []string{
		"", "einund", "zweiund", "dreiund", "vierund", "fünfund",
		"sechsund", "siebenund", "achtund", "neunund",
	}

	if n < 0 || n > 9 {
		panic("DigitToString1: invalid digit")
	}

	return dstrings[n]
}

// NumberToString returns a german string
// representation of the number n.
func NumberToString(n int) string {
	var s string

	if n%100 >= 0 && n%100 <= 19 {
		dstrings := []string{
			"null", "eins", "zwei", "drei", "vier", "fünf",
			"sechs", "sieben", "acht", "neun", "zehn",
			"elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn",
			"sechzehn", "siebzehn", "achtzehn", "neunzehn",
		}
		return DigitToString100(n/100) + dstrings[n%100]
	}

	digits := DigitList(n)

	switch len(digits) {
	case 0:
		s = "null"
	case 1:
		s = DigitToString1(digits[0])
	case 2:
		s = DigitToString1(digits[1]) + DigitToString10(digits[0])
	case 3:
		s = DigitToString100(digits[0]) + DigitToString1(digits[2]) + DigitToString10(digits[1])
	default:
		panic("NumberToString: invalid number")
	}

	return s
}

//	return s100 + s10 + s1

package numbers

import "fmt"

func ExampleDigitList() {
	digits := DigitList(1234567890)
	fmt.Println(digits)

	// Output:
	// [1 2 3 4 5 6 7 8 9 0]
}

func ExampleDigitToString100() {
	fmt.Println(DigitToString100(0))
	fmt.Println(DigitToString100(1))
	fmt.Println(DigitToString100(2))
	fmt.Println(DigitToString100(3))
	fmt.Println(DigitToString100(4))
	fmt.Println(DigitToString100(5))
	fmt.Println(DigitToString100(6))
	fmt.Println(DigitToString100(7))
	fmt.Println(DigitToString100(8))
	fmt.Println(DigitToString100(9))

	// Output:
	//
	// einhundert
	// zweihundert
	// dreihundert
	// vierhundert
	// fünfhundert
	// sechshundert
	// siebenhundert
	// achthundert
	// neunhundert
}

func ExampleDigitToString10() {
	fmt.Println(DigitToString10(0))
	fmt.Println(DigitToString10(1))
	fmt.Println(DigitToString10(2))
	fmt.Println(DigitToString10(3))
	fmt.Println(DigitToString10(4))
	fmt.Println(DigitToString10(5))
	fmt.Println(DigitToString10(6))
	fmt.Println(DigitToString10(7))
	fmt.Println(DigitToString10(8))
	fmt.Println(DigitToString10(9))

	// Output:
	//
	// zehn
	// zwanzig
	// dreißig
	// vierzig
	// fünfzig
	// sechzig
	// siebzig
	// achtzig
	// neunzig
}

func ExampleDigitToString1() {
	fmt.Println(DigitToString1(0))
	fmt.Println(DigitToString1(1))
	fmt.Println(DigitToString1(2))
	fmt.Println(DigitToString1(3))
	fmt.Println(DigitToString1(4))
	fmt.Println(DigitToString1(5))
	fmt.Println(DigitToString1(6))
	fmt.Println(DigitToString1(7))
	fmt.Println(DigitToString1(8))
	fmt.Println(DigitToString1(9))

	// Output:
	//
	// einund
	// zweiund
	// dreiund
	// vierund
	// fünfund
	// sechsund
	// siebenund
	// achtund
	// neunund
}

func ExampleNumberToString_Default() {
	fmt.Println(NumberToString(120))
	fmt.Println(NumberToString(121))
	fmt.Println(NumberToString(130))
	fmt.Println(NumberToString(131))
	fmt.Println(NumberToString(299))

	// Output:
	// einhundertzwanzig
	// einhunderteinundzwanzig
	// einhundertdreißig
	// einhunderteinunddreißig
	// zweihundertneunundneunzig
}

func ExampleNumberToString_special() {
	fmt.Println(NumberToString(1))
	fmt.Println(NumberToString(2))
	fmt.Println(NumberToString(3))
	fmt.Println(NumberToString(9))
	fmt.Println(NumberToString(22))
	fmt.Println(NumberToString(23))
	fmt.Println(NumberToString(99))
	fmt.Println(NumberToString(113))

	// Output:
	// eins
	// zwei
	// drei
	// neun
	// zweiundzwanzig
	// dreiundzwanzig
	// neunundneunzig
	// einhundertdreizehn
}

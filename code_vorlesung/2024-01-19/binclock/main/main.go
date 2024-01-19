package main

import (
	"fmt"

	"github.com/tel23a-inf/inf1-material/code_vorlesung/2024-01-19/binclock"
)

func main() {
	// fmt.Printf("Unix Time: %d\n", binclock.UnixTime())
	// fmt.Printf("Seconds: %d\n", binclock.Seconds())
	// fmt.Printf("Minutes: %d\n", binclock.Minutes())
	// fmt.Printf("Hours: %d\n", binclock.Hours())

	unix := binclock.UnixTime()
	//unix := int(time.Now().Unix())
	hours := (binclock.Hours(unix) + 1) % 24
	minutes := binclock.Minutes(unix)
	seconds := binclock.Seconds(unix)

	// Die eigentliche Binäruhr
	fmt.Println(binclock.BinaryToString(binclock.ToBinary(hours)))
	fmt.Println(binclock.BinaryToString(binclock.ToBinary(minutes)))
	fmt.Println(binclock.BinaryToString(binclock.ToBinary(seconds)))
}

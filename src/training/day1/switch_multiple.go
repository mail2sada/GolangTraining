package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	today := time.Now()
	var t  = today.Month()
 
	switch t {
	case time.January, time.February, time.March:
		fmt.Println("Q1 of the year...")
	case time.April, time.May, time.June:
		fmt.Println("Q2 of the year...")
	case time.July, time.August, time.September :
		fmt.Println("Q3 of the year.")
	default:
		fmt.Println("Q4 of the year")
	}
}
package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	today := time.Now()
 
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th..")
	case 10:
		fmt.Println("Today is 10th")
	case 15:
		fmt.Println("Today is 15th.")
	case 25:
		fmt.Println("Today is 25th")
	case 31:
		fmt.Println("Today is 31st.")
	default:
		fmt.Println("No information available for that day.")
	}
}
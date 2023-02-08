package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: More about date and time parsing...")
	//Parse YYYY-MM-DD
	timeT, _ := time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(timeT)

	//Parse YY-MM-DD
	timeT, _ = time.Parse("06-01-02", "20-01-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD
	timeT, _ = time.Parse("2006-Jan-02", "2020-Jan-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Jan-29 Wednesday 12:19:25")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Jan-29 Wednesday 12:19:25 AM IST +05:30")
	fmt.Println(timeT)
}

package testtime

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("testtime: init methos in testtime")
}



var A = time.Now().Add(time.Hour * 5)

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	fmt.Println("Demo: cpu core")

	fmt.Println(runtime.NumCPU())
	var a atomic.Int32 = 100
	atomic.AddInt32()
}

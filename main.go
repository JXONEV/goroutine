package main

import (
	"goroutine/agency"
	"time"
	"fmt"
)

func main()  {
    t1 := time.Now()
    agency.Syn()
    elapsed1 := time.Since(t1)
    fmt.Println("Time for synchronous is: ", elapsed1)
    t2 := time.Now()
    agency.Asy()
    elapsed2 := time.Since(t2)
    fmt.Println("Time for asynchronous is: ", elapsed2)
}

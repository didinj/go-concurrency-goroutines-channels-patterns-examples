package main

import (
    "fmt"
    "time"
)

var counter int

func increment() {
    counter++
}

func main() {
    for i := 0; i < 1000; i++ {
        go increment()
    }
    time.Sleep(1 * time.Second)
    fmt.Println("Counter:", counter)
}

package main

import (
    "fmt"
    "time"
)

func worker1(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "Worker 1 done"
}

func worker2(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "Worker 2 done"
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    go worker1(ch1)
    go worker2(ch2)

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}

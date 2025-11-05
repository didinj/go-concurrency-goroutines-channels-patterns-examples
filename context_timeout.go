package main

import (
    "context"
    "fmt"
    "time"
)

func slowOperation(ctx context.Context) {
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Println("Operation cancelled:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    go slowOperation(ctx)
    time.Sleep(3 * time.Second)
}

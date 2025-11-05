package concurrency

import (
    "sync"
    "testing"
    "time"
)

func slowTask() {
    time.Sleep(10 * time.Millisecond)
}

func runSequential() {
    for i := 0; i < 10; i++ {
        slowTask()
    }
}

func runConcurrent() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            slowTask()
        }()
    }
    wg.Wait()
}

func BenchmarkSequential(b *testing.B) {
    for i := 0; i < b.N; i++ {
        runSequential()
    }
}

func BenchmarkConcurrent(b *testing.B) {
    for i := 0; i < b.N; i++ {
        runConcurrent()
    }
}

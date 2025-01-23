```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        data := make(chan int, 10)

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("goroutine %d starting\n", i)
                        data <- i
                        fmt.Printf("goroutine %d sending data\n", i)
                }(i)
        }

        go func() {
                wg.Wait() // Wait for all goroutines to finish
                close(data) // Close the channel to signal no more data
        }()

        for v := range data {
                fmt.Printf("main goroutine receiving data %d\n", v)
        }

        fmt.Println("All goroutines finished")
}
```
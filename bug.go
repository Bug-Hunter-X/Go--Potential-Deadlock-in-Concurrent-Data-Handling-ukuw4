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
                for i := 0; i < 10; i++ {
                        v := <-data
                        fmt.Printf("main goroutine receiving data %d\n", v)
                }
        }()

        wg.Wait()
        fmt.Println("All goroutines finished")
}
```
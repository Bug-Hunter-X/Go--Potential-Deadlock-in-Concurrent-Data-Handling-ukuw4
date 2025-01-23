# Go: Potential Deadlock in Concurrent Data Handling

This repository demonstrates a common concurrency bug in Go: a potential deadlock scenario.  The provided `bug.go` file shows a program that can deadlock if the main goroutine consumes data from the channel faster than the worker goroutines can produce it. The `bugSolution.go` file illustrates a corrected version that avoids this deadlock.

## The Bug

The bug lies in the interaction between the main goroutine and the worker goroutines.  The main goroutine iterates to receive 10 items but does not wait for all goroutines to send data into the channel before consuming from it. This can lead to the main goroutine trying to read from an empty channel, causing a deadlock.

## The Solution

The solution uses a `sync.WaitGroup` to properly synchronize the main goroutine with the worker goroutines. The `WaitGroup` ensures the main goroutine waits for all worker goroutines to complete before exiting, preventing the deadlock.

## How to Run

1. Clone the repository
2. Navigate to the repository directory
3. Run `go run bug.go` to see the potential deadlock (may not always occur, depends on timing)
4. Run `go run bugSolution.go` to see the corrected, deadlock-free version

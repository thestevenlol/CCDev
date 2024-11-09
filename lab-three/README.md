# Barrier Synchronization Template in Go

Created by Jack Foley  
C00274246@setu.ie

This program demonstrates a barrier synchronization template using Go's concurrency primitives. It creates a set of goroutines that must wait at a barrier point before proceeding with the next part of their task, illustrating how barrier synchronization can be implemented with mutexes and channels.

## Overview

The program launches multiple goroutines, each executing two tasks:
- **Part A**: The initial task section.
- **Part B**: The final task section, which starts only once all goroutines have reached the barrier between the two parts.

### Features

- Barrier synchronization with `sync.Mutex` and unbuffered channels
- Concurrent execution and waiting of goroutines
- Simple example of mutual exclusion and controlled progression past a barrier

## Implementation Details

- `doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool)`: The main function for each goroutine. It performs Part A, waits at the barrier, and proceeds to Part B once all routines reach the barrier.
- `main()`: Sets up shared resources and launches the goroutines, waiting for all to complete before exiting.

### Important Notes

- **Barrier Mechanism**: An unbuffered channel is used to signal the completion of Part A by all goroutines, allowing them to synchronize before moving to Part B.
- **Synchronization**: The `sync.Mutex` ensures safe access to the shared `arrived` counter, while `sync.WaitGroup` waits for all goroutines to complete their execution.
- **Limitations**: The barrier is not fully implemented and may need refinement for other use cases or dynamic counts of goroutines.

## Use

To run the program:
```bash
go run main.go
```

## License
Concurrent Development SD4 © 2024 by Jack Foley is licensed under CC BY-NC-ND 4.0
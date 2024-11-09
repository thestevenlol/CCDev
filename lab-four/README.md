# Reusable Barrier Synchronization Template in Go

Created by Jack Foley  
C00274246@setu.ie

This program demonstrates a barrier synchronization template in Go. It allows multiple goroutines to reach a barrier point before they can proceed, simulating a scenario where all routines must reach a specific point before any of them can continue execution.

## Overview

The program initiates a set of goroutines to perform two-part tasks:
1. **Part A** simulates the initial phase of each goroutine's task.
2. **Part B** starts only once all goroutines have reached the barrier point between the two parts.

### Features

- Barrier synchronization using mutex and channels
- Concurrency control with goroutines
- Reset mechanism for multiple runs of the routine set

## Implementation Details

- `doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool)`: Each goroutine runs this function, where it completes Part A, waits at the barrier, and then completes Part B.
- `main()`: Manages the overall routine execution, including synchronization and resetting shared resources for each run.

### Important Notes

- **Barrier Mechanism**: An unbuffered channel acts as a signal to control when all goroutines proceed past the barrier.
- **Synchronization**: `sync.Mutex` and `sync.WaitGroup` ensure thread-safe operation and correct sequence of actions across goroutines.
- **Limitations**: The barrier is not implemented fully, requiring further refinement.

## Use

To run the program:
```bash
go run main.go
```

## License
Concurrent Development SD4 © 2024 by Jack Foley is licensed under CC BY-NC-ND 4.0
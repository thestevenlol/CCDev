# Producer-Consumer Problem Implementation in Go

Created by Jack Foley  
C00274246@setu.ie

This program implements the classic Producer-Consumer problem using Go's concurrency features. The implementation provides synchronization between producers and consumers using mutexes and channels, ensuring that no data races occur while items are produced and consumed.

## Overview
The program simulates a scenario with multiple producers and consumers where:
- Producers add items to a shared buffer (channel).
- Consumers remove items from the buffer.
- Mutexes and channels are used to manage access to the buffer and prevent race conditions.

### Features

- Concurrency with goroutines
- Synchronisation with `sync.Mutex` and `sync.WaitGroup`
- Buffered channel as a shared resource
- Safe concurrent production and consumption operations

## Implementation Details

The following are some of the important functions implemented in the program:

- `produce(wg *sync.WaitGroup, mutex *sync.Mutex, buf chan int, nProductions *int)`: Simulates producing an item and adding it to the buffer.
- `consume(wg *sync.WaitGroup, mutex *sync.Mutex, buf chan int, nConsumers *int)`: Simulates consuming an item from the buffer.
- `main()`: Initializes resources, starts producer and consumer goroutines, and waits for all to complete.

## Use

To run the program:
```bash
go run main.go
```

## License
Concurrent Development SD4 © 2024 by Jack Foley is licensed under CC BY-NC-ND 4.0
# Rendezvous Synchronization Example in Go

Created by Jack Foley

C00274246@setu.ie

This program shows a rendezvous synchronization pattern using Go's concurrency primitives. Here, several goroutines are started, running concurrently until they all reach a certain point—a rendezvous—waiting for all to arrive before continuing.

## Overview
The program starts 30 goroutines, each of which simulates a task with a random delay before arriving at a rendezvous point. The main function waits for each goroutine to signal its arrival at the rendezvous before allowing the program to continue.

### Features

- Rendezvous synchronization with channels
- Concurrency with goroutines
- Randomized timing for each goroutine to simulate varied processing times

## Implementation Details

- `rendezvousChannel`: A buffered channel to coordinate the rendezvous. Each goroutine signals its arrival by sending a message to this channel.
- **Goroutine Tasks**: Each goroutine sleeps for a random interval between 0.5 and 2 seconds to simulate different rendezvous arrivals.
- **Main Function**: Waits for signals from all goroutines, ensuring each has reached the rendezvous point before proceeding.

### Important Notes

- **Rendezvous Mechanism**: Each goroutine sends a signal to the `rendezvousChannel` upon arrival at the rendezvous point. 
The main function blocks until it has received `n` signals, indicating all goroutines have arrived. 
- **Concurrency**: This example uses Go's concurrency features to manage goroutine lifecycle and synchronization.

## Usage

To run the program:

```bash 
go run main.go 
``` 

## License Concurrent Development SD4 © 2024 by Jack Foley is licensed under CC BY-NC-ND 4.0
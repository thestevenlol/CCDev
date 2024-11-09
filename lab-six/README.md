# Dining Philosophers Problem Implementation in Go

Created by Jack Foley  
C00274246@setu.ie

This program implements the classic Dining Philosophers problem using Go's features for concurrency. The implementation provides fork synchronization by using channels and shows how to prevent deadlock by asymmetric resource acquisition.

## Overview
The program simulates 5 philosophers who sit around a circular table, each trying to alternate between thinking and eating. Between each philosopher is a fork, and a philosopher needs both his or her adjacent forks to eat.

### Features

- Concurrency with goroutines
- Deadlock avoidance by asymmetric fork pickup
- Random duration thinking and eating activities
- Using channel to synchronize the resources
- Thread-safe operation using sync.WaitGroup

## Implementation Details

The following are some of the important functions carried out by the program:

- `think(index int)`: Mock philosopher thinking for 0-4 seconds
- `eat(index int)`: Simulate eating with timing information
- `getForks(index int, forks map[int]chan bool)`: Asymmetric fork pick up implementation
- `putForks(index int, forks map[int]chan bool)`: Put down the fork after eating
- `doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool)`: Main philosopher lifecycle

## Use

To run the program:
```bash
go run main.go
```

## License
Concurrent Development SD4 © 2024 by Jack Foley is licensed under CC BY-NC-ND 4.0 
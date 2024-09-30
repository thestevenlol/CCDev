//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, mu *sync.Mutex, count *int, barrierSem *semaphore.Weighted, totalRoutines int, ctx context.Context) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// Barrier implementation
	mu.Lock()
	*count++
	if *count == totalRoutines {
		// All goroutines have reached the barrier; release them
		barrierSem.Release(1)
		mu.Unlock()
		err := barrierSem.Acquire(ctx, 1)
		if err != nil {
			return false
		}
	} else {
		mu.Unlock()
		// Wait for other goroutines to reach the barrier
		err := barrierSem.Acquire(ctx, 1)
		if err != nil {
			return false
		}
		barrierSem.Release(1)
	}

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	ctx := context.TODO()
	wg.Add(totalRoutines)
	var mu sync.Mutex
	var count int = 0
	// Initialize a semaphore with zero initial permits
	barrierSem := semaphore.NewWeighted(0)

	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &wg, &mu, &count, barrierSem, totalRoutines, ctx)
	}

	wg.Wait() // Wait for all goroutines to finish
}

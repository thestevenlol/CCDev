//Copyright (C) 2024 Mr. Jack Foley

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

//-------------------------------------------
// Author: Jack Foley (C00274246@setu.ie)
// Created on 07/10/2024
// Modified by: Jack Foley
// Issues: N/A
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

func doStuff(goNum int, arrived *int, finished *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	sharedLock.Lock()
	*arrived++
	if *arrived == max {
		sharedLock.Unlock()
		theChan <- true
		<-theChan
	} else {
		sharedLock.Unlock()
		<-theChan
		theChan <- true
	}
	sharedLock.Lock()
	*arrived--
	sharedLock.Unlock()
	fmt.Println("PartB", goNum)

	// Barrier mechanism for reusing the barrier
	// Requires resetting data by one of the threads once all the "PartB" have been printed

	sharedLock.Lock()
	*finished++
	if *finished == max {
		sharedLock.Unlock()
		theChan <- true
		<-theChan
	} else {
		sharedLock.Unlock()
		<-theChan
		theChan <- true
	}
	sharedLock.Lock()
	*finished--
	sharedLock.Unlock()

	if sharedLock.TryLock() {
		// reset the barrier
		*arrived = 0
		*finished = 0
		sharedLock.Unlock()
	}

	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	totalRuns := 10
	arrived := 0
	finished := 0

	var wg sync.WaitGroup
	var runs sync.WaitGroup
	var theLock sync.Mutex

	wg.Add(totalRoutines)
	runs.Add(totalRuns)
	theChan := make(chan bool) //use unbuffered channel in place of semaphore

	for _ = range totalRuns {
		for i := range totalRoutines { //create the go Routines here
			go doStuff(i, &finished, &arrived, totalRoutines, &wg, &theLock, theChan)
		}
		wg.Wait()
		runs.Done()
	}
	runs.Wait()
}

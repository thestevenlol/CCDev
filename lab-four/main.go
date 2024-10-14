//Barrier.go Template Code
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

//--------------------------------------------
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

func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
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
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	totalRuns := 10
	arrived := 0

	var wg sync.WaitGroup
	var runs sync.WaitGroup
	var theLock sync.Mutex

	wg.Add(totalRoutines)
	runs.Add(totalRuns)
	theChan := make(chan bool) //use unbuffered channel in place of semaphore

	for j := range totalRuns {
		for i := range totalRoutines { //create the go Routines here
			go doStuff(i, &arrived, totalRoutines, &wg, &theLock, theChan)
		}
		wg.Wait()
		fmt.Println("Done", j)    //wait for everyone to finish before next run
		wg.Add(totalRoutines)     //reset the wait group
		theChan = make(chan bool) //reset the channel
		theLock = sync.Mutex{}    //reset the lock
		arrived = 0               //reset the arrived counter
		runs.Done()
	}

	runs.Wait() //wait for all runs to finish
}

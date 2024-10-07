package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
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
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool)     //use unbuffered channel in place of semaphore
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
}

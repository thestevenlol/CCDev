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
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// think will have the philosopher think for a random amount of time
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// eat will have the philosopher eat for a random amount of time
func eat(index int) {
	fmt.Println("Phil: ", index, "is eating")
	now := time.Now()
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * (100 * time.Millisecond)) //wait random time amount
	elapsed := time.Since(now)
	fmt.Println("Phil: ", index, "done eating... took ", elapsed)
}

/*
getForks 	Will have the philosopher pick up the forks.

	If the index of the philosopher is even, they
	will pick up the fork to their right first. If
	the index of the philosopher is odd, they will
	pick up the fork to their left first.
*/
func getForks(index int, forks map[int]chan bool) {
	if index%2 == 0 {
		forks[index] <- true
		forks[(index+1)%5] <- true
	} else {
		forks[(index+1)%5] <- true
		forks[index] <- true
	}
}

func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	defer wg.Done()
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main

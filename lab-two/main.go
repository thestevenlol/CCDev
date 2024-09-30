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

// Created by: Jack Foley - C00274246
// Date: 26/09/2024

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 30 // Number of goroutines
	rendezvousChannel := make(chan struct{}, n)

	for i := 1; i <= n; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d started...\n", id)
			// sleep random time between 0.5 and 2 seconds
			time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)
			fmt.Printf("Goroutine %d has reached the rendezvous point!\n", id)
			rendezvousChannel <- struct{}{} // Send a signal to the channel
		}(i)
	}

	// Wait for all goroutines to reach the rendezvous point
	for i := 0; i < n; i++ {
		<-rendezvousChannel // Wait for n signals
	}
	fmt.Println("All goroutines have reached the rendezvous point. Proceeding!")
}
